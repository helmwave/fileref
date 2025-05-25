package internal

import (
	"fmt"
	"path/filepath"
)

type Render int

const (
	RenderCopy Render = iota
	RenderGomplate
	RenderSOPS
	RenderBase64
)

var extensionToRender = map[string]Render{
	".sops":   RenderSOPS,
	".base64": RenderBase64,
	".tpl":    RenderGomplate,
	".tmpl":   RenderGomplate,
}

func (r Render) String() string {
	return [...]string{"RENDER_COPY", "RENDER_GOMPLATE", "RENDER_SOPS", "RENDER_B64"}[r]
}

func (f *Config) render(r Render) error {
	switch r {
	case RenderCopy:
		return f.fetch()
	case RenderGomplate:
		return f.fetchAndRender()
	case RenderBase64:
		return f.fetchAndBase64()
	default:
		return fmt.Errorf("unknown render: %d", r)
	}
}

func (f *Config) Run() {
	r := f.getRenderTypeByExtension()
	err := f.render(r)
	if err != nil {
		f.l.Error(err.Error())
		panic(err)
	}
}

func (f *Config) getRenderTypeByExtension() Render {
	e := filepath.Ext(f.Dst)
	if r, ok := extensionToRender[e]; ok {
		return r
	}

	return RenderCopy
}
