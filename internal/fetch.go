package internal

import (
	"fmt"

	"github.com/hairyhenderson/gomplate/v4"
)

func (f *Config) inlineDatasource() string {
	return fmt.Sprintf("%s ds %q %s", f.G.LDelim, f.Name, f.G.RDelim)
}
func (f *Config) inlineTemplates() string {
	return fmt.Sprintf("%s template %q . %s", f.G.LDelim, f.Name, f.G.RDelim)
}

// Fetch fetches the template from the source
func (f *Config) Fetch() error {
	f.G.Input = f.inlineTemplates()
	f.G.Templates = f.selfMapDataSource()

	return f.Gomplate()
}

func (f *Config) selfMapDataSource() map[string]gomplate.DataSource {
	ds := make(map[string]gomplate.DataSource)
	ds[f.Name] = f.Src

	return ds
}
