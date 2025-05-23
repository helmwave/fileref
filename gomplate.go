package main

import (
	"github.com/hairyhenderson/gomplate/v4"
	"io"
	"os"
)

var DefaultGomplate = &gomplate.Config{
	RDelim:       "}}",
	LDelim:       "{{",
	Experimental: true,
}

func (f *Config) Gomplate() error {
	f.g.Stdout = f.File()
	defer f.File().Close()

	err := gomplate.Run(f.ctx, f.g)

	if err != nil && f.Strict {
		f.l.Error("Failed to gomplate: %v", err)
		panic(err)
	} else if err != nil {
		f.l.Warn("Failed to gomplate: %v", err)
	}

	return err
}

func (f *Config) Writer() io.Writer {
	return f.File()
}

func (f *Config) File() *os.File {
	file, err := os.Create(f.dst)
	if err != nil {
		f.l.Error("Failed to create file: %v", err)
		panic(err)
	}

	return file
}
