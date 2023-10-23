package config

import "fmt"

type Config interface {
	Download() error
}

type ConfigFS struct {
	Name string
}

func (c *ConfigFS) Download() error {
	if c.Name != "" {
		return fmt.Errorf("error")
	}
	return nil
}
