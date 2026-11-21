// Package output kirjoittaa työkalujen yhteisen JSON-kirjekuoren.
//
// Onnistunut tulos: {"schemaVersion":1,"data":...}
// Virhe:           {"schemaVersion":1,"error":{"code":"...","message":"..."}}
package output

import (
	"encoding/json"
	"io"
)

// SchemaVersion on kirjekuoren versio.
const SchemaVersion = 1

// Envelope on tulosteen kirjekuori.
type Envelope struct {
	SchemaVersion int    `json:"schemaVersion"`
	Data          any    `json:"data,omitempty"`
	Error         *Error `json:"error,omitempty"`
}

// Error on koneellisesti luettava virhe.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WriteData kirjoittaa onnistuneen tuloksen.
func WriteData(w io.Writer, data any) error {
	return write(w, Envelope{SchemaVersion: SchemaVersion, Data: data})
}

// WriteError kirjoittaa virheen.
func WriteError(w io.Writer, code, message string) error {
	return write(w, Envelope{SchemaVersion: SchemaVersion, Error: &Error{Code: code, Message: message}})
}

func write(w io.Writer, e Envelope) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(e)
}
