package config

import (
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

func getDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

//func DefaultConfig() *Config {
//
//}

func Load(path string) (*Config, error) {
	c := Config{}

	err := yaml.Unmarshal(os.ReadFile(path))
	return &c, err
}

//func (c *Config) Validate() error {
//
//}
