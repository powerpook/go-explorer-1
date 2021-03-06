package storage

import "github.com/centrifugal/gocent"

var publisher *gocent.Client

type CentrifugoConfig struct {
	Enable bool   `yaml:"enable"`
	Secret string `yaml:"secret"`
	URL    string `yaml:"url"`
	Socket string `yaml:"socket"`
	Key    string `yaml:"key"`
}

func (c *CentrifugoConfig) Initer() error {
	if c.Enable {
		publisher = gocent.New(gocent.Config{
			Addr: c.URL,
			Key:  c.Key,
		})
	}
	return nil
}

