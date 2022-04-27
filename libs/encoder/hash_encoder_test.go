package encoder

import (
	"testing"
)

func TestHashEncoder_EncodeAndDecode(t *testing.T) {

	tests := []struct {
		salt  string
		value int
	}{
		{
			salt:  "some salt",
			value: 1,
		},
		{
			salt:  "another salt",
			value: 4,
		},
		{
			salt:  "this salt",
			value: 25,
		},
		{
			salt:  "real big value",
			value: 999999999999999999,
		},
	}
	for _, tt := range tests {
		t.Run("should encode and decode given salt", func(t *testing.T) {
			e := NewHashEncoder(tt.salt)
			encoded, _ := e.Encode(tt.value)

			decoded, _ := e.Decode(encoded)

			if decoded != tt.value {
				t.Errorf("Decode value = %v, should return %v, got %v", encoded, tt.value, decoded)
			}
		})
	}
}
