// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-15 19:25:44.867410935 +0000 UTC

package optional

// Float64 is an optional float64
type Float64 struct {
	float64
	present bool
}

// EmptyFloat64 returns an empty optional.Float64
func EmptyFloat64() Float64 {
	return Float64{}
}

// OfFloat64 creates an optional.Float64 from a float64
func OfFloat64(f float64) Float64 {
	return Float64{float64: f, present: true}
}

// Set sets the float64 value
func (o *Float64) Set(f float64) {
	o.float64 = f
	o.present = true
}

// Float64 returns the float64 value
func (o *Float64) Float64() float64 {
	return o.float64
}

// Present returns whether or not the value is present
func (o *Float64) Present() bool {
	return o.present
}

// OrElse returns the float64 value or a default value if the value is not present
func (o *Float64) OrElse(f float64) float64 {
	if o.present {
		return o.float64
	}
	return f
}
