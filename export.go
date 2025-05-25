package main

import "github.com/helmwave/fileref/internal"

type Config = internal.Config

func New(name string) *Config {
	return internal.New(name)
}
