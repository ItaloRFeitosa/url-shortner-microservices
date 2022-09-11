package encoder

import (
	"github.com/italorfeitosa/url-shortner-mvp/libs/encoder"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/core"
)

type Encoder struct {
	encoder.HashEncoder
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) EncodeSlug(link *core.Link) error {
	encoded, err := e.Encode(int(link.ID))
	if err != nil {
		return err
	}
	link.Slug = encoded
	return nil
}
