package rand

import "testing"

func TestCryptoString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test 1", args{10}, "1234567", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CryptoString(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CryptoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
