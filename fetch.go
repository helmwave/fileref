package main

import (
	"fmt"
	"github.com/hairyhenderson/gomplate/v4"
)

func (f *Config) dsHack() string {
	return fmt.Sprintf("%s ds %q %s", f.g.LDelim, f.name, f.g.RDelim)
}
func (f *Config) tplHack() string {
	return fmt.Sprintf("%s template %q . %s", f.g.LDelim, f.name, f.g.RDelim)
}

// Fetch fetches the template from the source
func (f *Config) Fetch() error {
	ds := make(map[string]gomplate.DataSource)
	ds[f.name] = f.Src

	f.g.Input = f.tplHack()
	f.g.Templates = ds

	return f.Gomplate()
}
