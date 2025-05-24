package internal

import (
	"github.com/hairyhenderson/gomplate/v4"
)

type FileRefs map[string]*Config

// SetNames overrides the names of the FileRefs based on their keys.
func (rf FileRefs) SetNames() {
	for n, f := range rf {
		f.Name = n
	}
}

func (rf FileRefs) MapDataSources() map[string]gomplate.DataSource {
	m := make(map[string]gomplate.DataSource)
	for k, v := range rf {
		m[k] = v.Src
	}
	return m
}
