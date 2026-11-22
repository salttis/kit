// Package contextfile lukee Sontsan kontekstitiedoston (pakki/schemas/context.schema.json).
//
// Kontekstia ei ole pakko olla: Load palauttaa (nil, nil), kun SONTSA_CONTEXT_FILE
// puuttuu. Annettu mutta virheellinen konteksti on aina virhe.
package contextfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Ympäristömuuttujat, jotka Sontsa asettaa lapselle.
const (
	EnvContextFile     = "SONTSA_CONTEXT_FILE"
	EnvContractVersion = "SONTSA_CONTRACT_VERSION"
)

// SupportedContract on tuettu sopimusversio.
const SupportedContract = 1

// Context on yhden projektin ajokonteksti. Se sisältää viitteitä salaisuuksiin,
// ei salaisuuksien arvoja.
type Context struct {
	SchemaVersion   int               `json:"schemaVersion"`
	ContractVersion int               `json:"contractVersion"`
	ProjectID       string            `json:"projectId"`
	InvocationID    string            `json:"invocationId"`
	Environment     Environment       `json:"environment"`
	Config          map[string]any    `json:"config"`
	SecretRefs      map[string]string `json:"secretRefs"`
}

// Environment on valittu ympäristö.
type Environment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// Load lukee kontekstin ympäristömuuttujan osoittamasta tiedostosta.
func Load() (*Context, error) {
	path := os.Getenv(EnvContextFile)
	if path == "" {
		return nil, nil
	}
	if v := os.Getenv(EnvContractVersion); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n != SupportedContract {
			return nil, fmt.Errorf("%s=%q: tuettu sopimusversio on %d", EnvContractVersion, v, SupportedContract)
		}
	}
	return LoadFile(path)
}

// LoadFile lukee ja validoi kontekstitiedoston.
func LoadFile(path string) (*Context, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("kontekstitiedosto: %w", err)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	var c Context
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("kontekstitiedosto %s: %w", path, err)
	}
	if err := c.validate(); err != nil {
		return nil, fmt.Errorf("kontekstitiedosto %s: %w", path, err)
	}
	return &c, nil
}

func (c *Context) validate() error {
	if c.SchemaVersion != 1 {
		return fmt.Errorf("schemaVersion %d ei ole tuettu", c.SchemaVersion)
	}
	if c.ContractVersion != SupportedContract {
		return fmt.Errorf("contractVersion %d ei ole tuettu", c.ContractVersion)
	}
	if c.ProjectID == "" || c.InvocationID == "" {
		return errors.New("projectId ja invocationId vaaditaan")
	}
	if c.Environment.ID == "" || c.Environment.Name == "" || c.Environment.Kind == "" {
		return errors.New("environment.id, name ja kind vaaditaan")
	}
	switch c.Environment.Kind {
	case "place", "lab", "remote":
	default:
		return fmt.Errorf("environment.kind %q ei ole place, lab tai remote", c.Environment.Kind)
	}
	return nil
}
