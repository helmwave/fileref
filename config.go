package main

import (
	"context"
	"github.com/hairyhenderson/gomplate/v4"
	"log/slog"
)

func New(name string) *Config {
	return &Config{
		ctx:    context.Background(),
		Strict: true,
		name:   name,
		g:      DefaultGomplate,
		l:      DefaultLogger.With("FileRef", name),
	}
}

type Config struct {
	ctx context.Context
	g   *gomplate.Config
	l   *slog.Logger

	name   string
	Src    gomplate.DataSource
	dst    string
	Strict bool
}

func (f *Config) GetName() string {
	return f.name
}

func (f *Config) IsStrict() bool {
	return f.Strict
}
