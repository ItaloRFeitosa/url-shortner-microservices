package encoder

import "github.com/speps/go-hashids/v2"

type HashEncoder struct {
	h *hashids.HashID
}

func NewHashEncoder(salt string) *HashEncoder {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 4
	h, _ := hashids.NewWithData(hd)
	return &HashEncoder{h}
}

func (e *HashEncoder) Encode(v int) (string, error) {
	return e.h.Encode([]int{v})
}

func (e *HashEncoder) Decode(v string) (int, error) {
	decoded, err := e.h.DecodeWithError(v)

	if err != nil {
		return 0, err
	}

	return decoded[0], nil
}
