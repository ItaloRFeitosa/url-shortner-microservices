package map_helper

import (
	"github.com/mitchellh/mapstructure"
)

func Decode[T any](v any) (T, error) {
	var decoded T

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &decoded,
		TagName:  "json",
	})

	err := decoder.Decode(v)

	if err != nil {
		return decoded, err
	}

	return decoded, nil
}
