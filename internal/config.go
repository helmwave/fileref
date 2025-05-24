package internal

import (
	"context"
	"log/slog"

	"github.com/hairyhenderson/gomplate/v4"
)

func New(name string) *Config {
	return &Config{
		Ctx:    context.Background(),
		Strict: true,
		Name:   name,
		G:      DefaultGomplate,
		L:      DefaultLogger.With("FileRef", name),
	}
}

type Config struct {
	Ctx context.Context
	G   *gomplate.Config
	L   *slog.Logger

	Name   string
	Src    gomplate.DataSource
	Dst    string
	Strict bool
}

func (f *Config) String() string {
	return f.Name + " -> " + f.Dst
}
