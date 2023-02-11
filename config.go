package fiberlogrus

import "github.com/sirupsen/logrus"

// Config defines the config for middleware
type Config struct {
	Logger *logrus.Logger
	Tags   []string
}

// ConfigDefault is the default config
var ConfigDefault Config = Config{
	Logger: nil,
	Tags:   []string{TagMethod},
}
