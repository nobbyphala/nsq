package main

import (
	"testing"

	"github.com/mreiferson/go-options"
	"github.com/nobbyphala/nsq/internal/lg"
	"github.com/nobbyphala/nsq/internal/test"
	"github.com/nobbyphala/nsq/nsqadmin"
)

func TestConfigFlagParsing(t *testing.T) {
	opts := nsqadmin.NewOptions()
	opts.Logger = test.NewTestLogger(t)

	flagSet := nsqadminFlagSet(opts)
	flagSet.Parse([]string{})

	cfg := config{"log_level": "debug"}
	cfg.Validate()

	options.Resolve(opts, flagSet, cfg)
	if opts.LogLevel != lg.DEBUG {
		t.Fatalf("log level: want debug, got %s", opts.LogLevel.String())
	}
}
