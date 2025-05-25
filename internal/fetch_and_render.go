package internal

import "fmt"

// inlineGomplate is a hack for self-templating
func (f *Config) inlineGomplate() string {
	return fmt.Sprintf("%s template %q . %s", f.g.LDelim, f.Name, f.g.RDelim)
}

func (f *Config) fetchAndRender() error {
	f.g.Input = f.inlineGomplate()
	f.g.Templates = f.selfMapDataSource()

	return f.runGomplate()
}
