// Package config tarjoaa työkalujen yhteiset asetuspolut ja TOML-luvun.
//
// Projektin asetukset ovat os.UserConfigDir()/<nimi> ja välimuisti
// os.UserCacheDir()/<nimi>. Mikään funktio ei luo hakemistoja lukiessaan.
package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Dir palauttaa työkalun asetushakemiston (ei luo sitä).
func Dir(tool string) (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, tool), nil
}

// CacheDir palauttaa työkalun välimuistihakemiston (ei luo sitä).
func CacheDir(tool string) (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, tool), nil
}

// LoadTOML lukee TOML-tiedoston kohteeseen v. Tuntematon avain on virhe.
// Puuttuva tiedosto palauttaa fs.ErrNotExist-virheen (errors.Is toimii),
// jotta kutsuja päättää, ovatko oletukset sallittuja.
func LoadTOML(path string, v any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	md, err := toml.NewDecoder(bytes.NewReader(raw)).Decode(v)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	if un := md.Undecoded(); len(un) > 0 {
		return fmt.Errorf("%s: tuntematon avain %s", path, un[0])
	}
	return nil
}

// Missing kertoo, onko virhe puuttuvan tiedoston virhe.
func Missing(err error) bool { return errors.Is(err, fs.ErrNotExist) }
