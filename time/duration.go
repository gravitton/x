package time

import "time"

// Duration is a time.Duration that marshals to/from its string representation (e.g. "200ms").
type Duration time.Duration

// MarshalText encodes the duration as its string representation, e.g. "200ms".
func (d Duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

// UnmarshalText decodes a duration from its string representation, e.g. "200ms".
func (d *Duration) UnmarshalText(data []byte) error {
	v, err := time.ParseDuration(string(data))
	*d = Duration(v)
	return err
}

// String returns the duration in a human-readable form, e.g. "200ms".
func (d Duration) String() string {
	return time.Duration(d).String()
}

// Equal reports whether d equals the given time.Duration.
func (d Duration) Equal(other time.Duration) bool {
	return time.Duration(d) == other
}
