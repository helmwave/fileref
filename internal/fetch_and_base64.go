package internal

import (
	"fmt"
	"github.com/hairyhenderson/gomplate/v4"
	"time"
)

// inlineGomplate is a hack for self-templating
func (f *Config) inlineBase64() string {
	s := fmt.Sprintf("%s (ds %q | b64 ) %s", f.g.LDelim, f.Name, f.g.RDelim)
	return s
}

func (f *Config) fetchAndBase64() error {
	f.disableGomplateProcessing()
	f.g.Input = f.inlineBase64()
	f.g.DataSources = f.selfMapDataSource()
	f.setPluginBase64()

	return f.runGomplate()
}

func (f *Config) setPluginBase64() {
	if f.g.Plugins == nil {
		f.g.Plugins = make(map[string]gomplate.PluginConfig)
	}

	f.g.Plugins["b64"] = gomplate.PluginConfig{
		Cmd:     "/usr/bin/base64", // todo: How to Pass the path to base64?
		Args:    []string{"-d"},
		Pipe:    true,
		Timeout: 2 * time.Minute,
	}
}
