package config

import (
	"flag"
)

type Options struct {
	RunAddress string
}

func InitOptions() Options {
	options := new(Options)
	flag.StringVar(&options.RunAddress, "a", ":8080", "address and port to run server")

	flag.Parse()
	return *options
}
