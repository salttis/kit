// Package ui sisältää työkalujen yhteisen Lip Gloss -teeman.
//
// Teema käyttää adaptiivisia värejä, joten se toimii vaalealla ja tummalla
// pohjalla. TUI-mallit omistaa kukin työkalu; tähän kuuluu vain yhteinen ilme.
package ui

import "github.com/charmbracelet/lipgloss"

// Theme on yhteinen tyyliryhmä.
type Theme struct {
	Title   lipgloss.Style
	Subtle  lipgloss.Style
	OK      lipgloss.Style
	Warn    lipgloss.Style
	Err     lipgloss.Style
	Help    lipgloss.Style
	Focused lipgloss.Style
}

// Default palauttaa oletusteeman.
func Default() Theme {
	c := func(light, dark string) lipgloss.AdaptiveColor {
		return lipgloss.AdaptiveColor{Light: light, Dark: dark}
	}
	return Theme{
		Title:   lipgloss.NewStyle().Bold(true).Foreground(c("#5A56E0", "#9D9BFF")),
		Subtle:  lipgloss.NewStyle().Foreground(c("#6B6B6B", "#8A8A8A")),
		OK:      lipgloss.NewStyle().Foreground(c("#1E7E34", "#5FD07A")),
		Warn:    lipgloss.NewStyle().Foreground(c("#9A6700", "#E3B341")),
		Err:     lipgloss.NewStyle().Foreground(c("#B3261E", "#FF7B72")),
		Help:    lipgloss.NewStyle().Foreground(c("#6B6B6B", "#8A8A8A")),
		Focused: lipgloss.NewStyle().Bold(true).Foreground(c("#5A56E0", "#9D9BFF")),
	}
}

// State palauttaa tilan tyylin: "ok", "warn", "error"; muu on Subtle.
func (t Theme) State(state string) lipgloss.Style {
	switch state {
	case "ok":
		return t.OK
	case "warn":
		return t.Warn
	case "error":
		return t.Err
	}
	return t.Subtle
}
