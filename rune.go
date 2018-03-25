// Code generated by go generate
// This file was generated by robots at 2018-03-25 18:09:24.186098998 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Rune is an optional rune
type Rune struct {
	value *rune
}

// NewRune creates a optional.Rune from a rune
func NewRune(v rune) Rune {
	return Rune{&v}
}

// Set sets the rune value
func (r Rune) Set(v rune) {
	r.value = &v
}

// Get returns the rune value or an error if not present
func (r Rune) Get() (rune, error) {
	if !r.Present() {
		return *r.value, errors.New("value not present")
	}
	return *r.value, nil
}

// Present returns whether or not the value is present
func (r Rune) Present() bool {
	return r.value != nil
}

// OrElse returns the rune value or a default value if the value is not present
func (r Rune) OrElse(v rune) rune {
	if r.Present() {
		return *r.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (r Rune) If(fn func(rune)) {
	if r.Present() {
		fn(*r.value)
	}
}

func (r Rune) MarshalJSON() ([]byte, error) {
	if r.Present() {
		return json.Marshal(r.value)
	}
	return nil, nil
}

func (r *Rune) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		r.value = nil
		return nil
	}

	r.value = new(rune)
	return json.Unmarshal(data, r.value)
}
