// Package logx rakentaa työkalujen yhteisen log/slog-lokittajan.
//
// Oletus on stderr ja taso warn; verbose nostaa tason debugiin. Kentät:
// aika (UTC), taso, project, version ja invocationId. Salaisuuksia ei lokiteta.
package logx

import (
	"io"
	"log/slog"
	"time"
)

// Options määrittää lokittajan.
type Options struct {
	Project      string
	Version      string
	InvocationID string
	Verbose      bool
	JSON         bool
}

// New palauttaa lokittajan, joka kirjoittaa w:hen (yleensä os.Stderr, TUI:ssa
// tiedosto tai io.Discard, jotta näyttö ei sotkeudu).
func New(w io.Writer, o Options) *slog.Logger {
	level := slog.LevelWarn
	if o.Verbose {
		level = slog.LevelDebug
	}
	ho := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().UTC().Format(time.RFC3339))
			}
			return a
		},
	}
	var h slog.Handler
	if o.JSON {
		h = slog.NewJSONHandler(w, ho)
	} else {
		h = slog.NewTextHandler(w, ho)
	}
	l := slog.New(h).With("project", o.Project, "version", o.Version)
	if o.InvocationID != "" {
		l = l.With("invocationId", o.InvocationID)
	}
	return l
}
