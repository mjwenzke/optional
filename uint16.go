// Code generated by go generate
// This file was generated by robots at 2018-03-25 18:09:22.241969327 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Uint16 is an optional uint16
type Uint16 struct {
	value *uint16
}

// NewUint16 creates a optional.Uint16 from a uint16
func NewUint16(v uint16) Uint16 {
	return Uint16{&v}
}

// Set sets the uint16 value
func (u Uint16) Set(v uint16) {
	u.value = &v
}

// Get returns the uint16 value or an error if not present
func (u Uint16) Get() (uint16, error) {
	if !u.Present() {
		return *u.value, errors.New("value not present")
	}
	return *u.value, nil
}

// Present returns whether or not the value is present
func (u Uint16) Present() bool {
	return u.value != nil
}

// OrElse returns the uint16 value or a default value if the value is not present
func (u Uint16) OrElse(v uint16) uint16 {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (u Uint16) If(fn func(uint16)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	return nil, nil
}

func (u *Uint16) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		u.value = nil
		return nil
	}

	u.value = new(uint16)
	return json.Unmarshal(data, u.value)
}
