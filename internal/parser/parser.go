package parser

import (
	"github.com/css-scripter/healthi/internal/types"
	"gopkg.in/yaml.v3"
)

func ParseConfig(target *types.Config, file []byte) error {
	return yaml.Unmarshal(file, &target)
}
