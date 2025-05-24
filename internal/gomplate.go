package internal

import (
	"io"
	"os"

	"github.com/hairyhenderson/gomplate/v4"
)

var DefaultGomplate = &gomplate.Config{
	RDelim:       "}}",
	LDelim:       "{{",
	Experimental: true,
}

func (f *Config) Gomplate() error {
	f.G.Stdout = f.File()
	defer f.File().Close()

	err := gomplate.Run(f.Ctx, f.G)

	if err != nil && f.Strict {
		f.L.Error("Failed to gomplate: %v", err)
		panic(err)
	} else if err != nil {
		f.L.Warn("Failed to gomplate: %v", err)
	}

	return err
}

func (f *Config) Writer() io.Writer {
	return f.File()
}

func (f *Config) File() *os.File {
	file, err := os.Create(f.Dst)
	if err != nil {
		f.L.Error("Failed to create file: %v", err)
		panic(err)
	}

	return file
}
