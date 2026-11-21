// Package buildinfo kertoo ohjelman version ja build-tiedot.
//
// Työkalu asettaa arvot build-vaiheessa:
//
//	-ldflags "-X github.com/salttis/kit/buildinfo.Version=1.2.0 -X ...Commit=abc -X ...BuildTime=2026-11-22T10:00:00Z"
//
// Puuttuva arvo on "unknown".
package buildinfo

import (
	"runtime"
	"runtime/debug"
)

// ContractVersion on JS-prosessisopimuksen versio (js/docs/02).
const ContractVersion = 1

// Ldflags-arvot.
var (
	Version   = ""
	Commit    = ""
	BuildTime = ""
)

// Info on `version --output json` -tuloksen data-olio.
type Info struct {
	Version         string `json:"version"`
	Commit          string `json:"commit"`
	BuildTime       string `json:"buildTime"`
	GoVersion       string `json:"goVersion"`
	ContractVersion int    `json:"contractVersion"`
}

// Get palauttaa version tiedot. Jos ldflags puuttuvat, commit luetaan Go:n
// build-metatiedoista, kun se on saatavilla.
func Get() Info {
	info := Info{
		Version:         orUnknown(Version),
		Commit:          orUnknown(Commit),
		BuildTime:       orUnknown(BuildTime),
		GoVersion:       runtime.Version(),
		ContractVersion: ContractVersion,
	}
	if Commit == "" {
		if bi, ok := debug.ReadBuildInfo(); ok {
			for _, s := range bi.Settings {
				if s.Key == "vcs.revision" && s.Value != "" {
					info.Commit = s.Value
				}
			}
		}
	}
	return info
}

func orUnknown(s string) string {
	if s == "" {
		return "unknown"
	}
	return s
}
