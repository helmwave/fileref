package internal

import (
	"fmt"
)

// inlineDataSource is a hack for self-datasourcing within gomplate
func (f *Config) inlineDataSource() string {
	return fmt.Sprintf("%s ds %q %s", f.g.LDelim, f.Name, f.g.RDelim)
}

// fetch fetches the template from the source
func (f *Config) fetch() error {
	f.disableGomplateProcessing()
	f.g.Input = f.inlineDataSource()
	f.g.DataSources = f.selfMapDataSource()

	return f.runGomplate()
}

func (f *Config) disableGomplateProcessing() {
	//f.Src.Header.Set("Content-Type", "text/plain")
	q := f.Src.URL.Query()
	q.Set("type", "text/plain")
	f.Src.URL.RawQuery = q.Encode()

	f.l.Debug("disabling runGomplate processing", "source", f.Src.URL.String())
}
