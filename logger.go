package fiberlogrus

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// getLogrusFields calls FuncTag functions on matching keys
func getLogrusFields(ftm map[string]FuncTag, c *fiber.Ctx, d *data) log.Fields {
	f := make(log.Fields)
	for k, ft := range ftm {
		f[k] = ft(c, d)
	}
	return f
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	var cfg Config
	if len(config) == 0 {
		cfg = ConfigDefault
	} else {
		cfg = config[0]
	}
	d := new(data)
	// Set PID once
	d.pid = os.Getpid()
	ftm := getFuncTagMap(cfg, d)

	return func(c *fiber.Ctx) error {
		d.start = time.Now()

		err := c.Next()

		d.end = time.Now()
		var logEntry *log.Entry
		switch cfg.Logger {
		case nil:
			logEntry = log.WithFields(getLogrusFields(ftm, c, d))
		default:
			logEntry = cfg.Logger.WithFields(getLogrusFields(ftm, c, d))
		}

		if err == nil {
			logEntry.Info()
		} else {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			logEntry.WithField("status", code).Error(err)
		}

		return err
	}
}
