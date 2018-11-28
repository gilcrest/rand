package rand

import (
	mrand "math/rand"
	"time"
)

// Charset defines a character set allowed for the random string
type Charset uint8

// Different Character Set values
const (
	Default Charset = iota
)

func (c Charset) String() string {
	switch c {
	case Default:
		return "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	return "unknown Character Set"
}

var seededRand = mrand.New(mrand.NewSource(time.Now().UnixNano()))

// MathStringWCharset - given a character set, returns a randomly generated
// string using the math/rand package. This should be used when performance
// is needed and there are less concerns for security.
func MathStringWCharset(length int, c Charset) string {
	charset := c.String()
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// MathString returns a randomly generated string using the math/rand
// package and the Default character set. This should be used when
// performance is needed and there are less concerns for security.
func MathString(length int) string {
	return MathStringWCharset(length, Default)
}
