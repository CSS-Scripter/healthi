package main

import (
	"fmt"
	"os"

	"github.com/css-scripter/healthi/internal/parser"
	"github.com/css-scripter/healthi/internal/scheduler"
	"github.com/css-scripter/healthi/internal/types"
	"github.com/css-scripter/healthi/pkg/fn"
)

func main() {
	config := readConfig("./healthi.yml")
	initHealthi(config)
}

func readConfig(fileName string) types.Config {
	var data []byte
	var config types.Config
	var err error

	flow := fn.Flow{Slice: fn.Slice[func() error]{
		func() error { data, err = os.ReadFile(fileName); return err },
		func() error { return parser.ParseConfig(&config, data) },
	}}

	err = flow.Run()
	if err != nil {
		panic(err)
	}
	return config
}

func initHealthi(config types.Config) {
	errorsChan := make(chan error)
	config.Services.ForEach(func(cs types.ConfigService) bool {
		go scheduler.RunTask(errorsChan, cs.Run, cs.Interval)
		return true
	})

	for err := range errorsChan {
		fmt.Println(err.Error())
	}
}
