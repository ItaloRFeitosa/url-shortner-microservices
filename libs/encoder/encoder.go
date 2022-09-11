package encoder

type Encoder interface {
	Encode(v int) (string, error)
}

type Decoder interface {
	Decode(v string) (int, error)
}

type Codec interface {
	Encoder
	Decoder
}
