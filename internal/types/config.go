package types

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/css-scripter/healthi/pkg/fn"
)

const (
	FailureStrategyAll    = "all"
	FailureStrategySingle = "single"
)

type Config struct {
	Services fn.Slice[ConfigService] `yaml:"services"`
}

type ConfigService struct {
	Name            string                 `yaml:"name"`
	FailureStrategy string                 `yaml:"failure_strategy"`
	FailureCount    int                    `yaml:"failure_count"`
	Interval        int                    `yaml:"interval"`
	Scrape          fn.Slice[ConfigScrape] `yaml:"scrape"`
	Alerting        fn.Slice[ConfigAlert]  `yaml:"alerting"`
}

type ConfigScrape struct {
	Name   string           `yaml:"name"`
	Urls   fn.Slice[string] `yaml:"urls"`
	Method string           `yaml:"method"`
	Status int              `yaml:"status"`
}

type ConfigAlert struct {
	Name     string                 `yaml:"name"`
	Type     string                 `yaml:"type"`
	Settings fn.Map[string, string] `yaml:"settings"`
}

func (cs ConfigService) Run() error {
	failures := fn.Slice[string]{}

	cs.Scrape.ForEach(func(scrape ConfigScrape) bool {
		err := scrape.Run(cs.FailureStrategy)
		if err != nil {
			failures = append(failures, scrape.Name)
		}
		return true
	})

	if cs.hasFailed(failures) {
		return errors.New(fmt.Sprintf("[%s] Failure; Failed checks: [%s]", cs.Name, strings.Join(failures, ", ")))
	}

	return nil
}

func (cs ConfigService) hasFailed(failures fn.Slice[string]) bool {
	return (cs.FailureStrategy == FailureStrategyAll && len(failures) == len(cs.Scrape)) ||
		(cs.FailureStrategy == FailureStrategySingle && len(cs.Scrape) > 0)
}

func (cs ConfigScrape) Run(failureStrategy string) error {
	var err error
	allFailed := true
	cs.Urls.ForEach(func(s string) bool {
		var res *http.Response
		res, err = http.Get(s)
		if err != nil || res.StatusCode != cs.Status {
			return failureStrategy == FailureStrategyAll
		}
		allFailed = false
		return true
	})

	if failureStrategy == FailureStrategyAll {
		if allFailed {
			return errors.New("all scrapes failed")
		} else {
			return nil
		}
	}

	return err
}
