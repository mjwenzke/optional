// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-15 19:25:42.408641543 +0000 UTC

package optional

// Int64 is an optional int64
type Int64 struct {
	int64
	present bool
}

// EmptyInt64 returns an empty optional.Int64
func EmptyInt64() Int64 {
	return Int64{}
}

// OfInt64 creates an optional.Int64 from a int64
func OfInt64(i int64) Int64 {
	return Int64{int64: i, present: true}
}

// Set sets the int64 value
func (o *Int64) Set(i int64) {
	o.int64 = i
	o.present = true
}

// Int64 returns the int64 value
func (o *Int64) Int64() int64 {
	return o.int64
}

// Present returns whether or not the value is present
func (o *Int64) Present() bool {
	return o.present
}

// OrElse returns the int64 value or a default value if the value is not present
func (o *Int64) OrElse(i int64) int64 {
	if o.present {
		return o.int64
	}
	return i
}
