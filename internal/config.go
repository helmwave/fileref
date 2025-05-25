package internal

import (
	"context"
	"github.com/hairyhenderson/gomplate/v4"
	"log/slog"
	"net/http"
	"net/url"
	"path/filepath"
)

func New(name string) *Config {
	return &Config{
		ctx:    context.Background(),
		Strict: true,
		Name:   name,
		g:      DefaultGomplate,
		l:      DefaultLogger.With("FileRef", name),
	}
}

type Config struct {
	ctx context.Context
	g   *gomplate.Config
	l   *slog.Logger

	Name   string
	Src    gomplate.DataSource
	Dst    string
	Strict bool
}

func (f *Config) String() string {
	return f.Name + " -> " + f.Dst
}

func (f *Config) selfMapDataSource() map[string]gomplate.DataSource {
	ds := make(map[string]gomplate.DataSource)
	ds[f.Name] = f.Src

	return ds
}

func (f *Config) SetSrc(uri string) {
	u, err := url.Parse(uri)
	if err != nil {
		f.l.Error("failed to parse source URL", "uri", uri, "error", err)
		panic(err)
	}

	f.Src = gomplate.DataSource{
		URL:    u,
		Header: make(http.Header),
	}

	f.Dst = f.Name
}

const defaultSrcExtension = ".yml"

func (f *Config) getSrcExtension() string {
	e := filepath.Ext(f.Src.URL.Path)
	if "" == e {
		return defaultSrcExtension
	}

	return e
}

func (f *Config) getNameExtension() string {
	e := filepath.Ext(f.Name)
	if "" == e {
		return defaultSrcExtension
	}

	return e
}

func (f *Config) SetCtx(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	f.ctx = ctx
}
