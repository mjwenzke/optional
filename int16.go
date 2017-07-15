// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-15 19:25:41.945778486 +0000 UTC

package optional

// Int16 is an optional int16
type Int16 struct {
	int16
	present bool
}

// EmptyInt16 returns an empty optional.Int16
func EmptyInt16() Int16 {
	return Int16{}
}

// OfInt16 creates an optional.Int16 from a int16
func OfInt16(i int16) Int16 {
	return Int16{int16: i, present: true}
}

// Set sets the int16 value
func (o *Int16) Set(i int16) {
	o.int16 = i
	o.present = true
}

// Int16 returns the int16 value
func (o *Int16) Int16() int16 {
	return o.int16
}

// Present returns whether or not the value is present
func (o *Int16) Present() bool {
	return o.present
}

// OrElse returns the int16 value or a default value if the value is not present
func (o *Int16) OrElse(i int16) int16 {
	if o.present {
		return o.int16
	}
	return i
}
