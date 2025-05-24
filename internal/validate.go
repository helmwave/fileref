package internal

import "github.com/pkg/errors"

func (f *Config) Validate() error {
	if f.Src.URL.String() == f.Dst {
		return errors.New("the source and destination are the same")
	}

	return nil
}
