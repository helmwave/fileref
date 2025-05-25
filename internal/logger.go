package internal

import (
	"log/slog"
	"os"
)

var DefaultLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

func init() {
	slog.SetDefault(DefaultLogger)
	slog.Info("DefaultLogger initialized")
}

func (f *Config) SetLogger(l *slog.Logger) {
	if l == nil {
		l = slog.Default()
	}

	f.l = l.With("FileRef", f.Name)
}
