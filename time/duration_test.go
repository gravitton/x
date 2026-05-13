package time

import (
	"testing"
	"time"

	"github.com/gravitton/assert"
)

func TestDuration_Equal(t *testing.T) {
	d := Duration(time.Millisecond * 200)

	assert.True(t, d.Equal(time.Millisecond*200))
	assert.False(t, d.Equal(time.Millisecond*100))
}

func TestDuration_String(t *testing.T) {
	cases := []struct {
		d    Duration
		want string
	}{
		{Duration(time.Millisecond * 200), "200ms"},
		{Duration(time.Second * 5), "5s"},
		{Duration(time.Minute * 3), "3m0s"},
		{Duration(time.Hour * 2), "2h0m0s"},
		{Duration(time.Hour + 3*time.Minute + 4*time.Second), "1h3m4s"},
		{Duration(2*time.Hour + 30*time.Minute + 500*time.Millisecond), "2h30m0.5s"},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, c.d.String())
	}
}

func TestDuration_MarshalText(t *testing.T) {
	cases := []struct {
		d    Duration
		want string
	}{
		{Duration(time.Millisecond * 200), "200ms"},
		{Duration(time.Second * 5), "5s"},
		{Duration(time.Minute * 3), "3m0s"},
		{Duration(time.Hour * 2), "2h0m0s"},
		{Duration(time.Hour + 3*time.Minute + 4*time.Second), "1h3m4s"},
		{Duration(2*time.Hour + 30*time.Minute + 500*time.Millisecond), "2h30m0.5s"},
	}
	for _, c := range cases {
		b, err := c.d.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, c.want, string(b))
	}
}

func TestDuration_UnmarshalText(t *testing.T) {
	cases := []struct {
		input string
		want  time.Duration
	}{
		{"200ms", time.Millisecond * 200},
		{"5s", time.Second * 5},
		{"3m0s", time.Minute * 3},
		{"2h0m0s", time.Hour * 2},
		{"1h3m4s", time.Hour + 3*time.Minute + 4*time.Second},
		{"2h30m0.5s", 2*time.Hour + 30*time.Minute + 500*time.Millisecond},
	}
	for _, c := range cases {
		var d Duration
		err := d.UnmarshalText([]byte(c.input))
		assert.NoError(t, err)
		assert.True(t, d.Equal(c.want))
	}
}

func TestDuration_UnmarshalText_Invalid(t *testing.T) {
	var d Duration
	err := d.UnmarshalText([]byte("invalid"))

	assert.Error(t, err)
}
