package rand

import (
	"testing"
)

func TestMathString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{20}, "123457"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MathString(tt.args.length); got != tt.want {
				t.Errorf("MathString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathStringWCharset(t *testing.T) {
	type args struct {
		length int
		c      Charset
	}

	charset := Default

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{30, charset}, "1234567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MathStringWCharset(tt.args.length, tt.args.c); got != tt.want {
				t.Errorf("MathStringWCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}
