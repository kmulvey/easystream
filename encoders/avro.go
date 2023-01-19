package encoders

import "github.com/linkedin/goavro/v2"

type AvroEncoder struct {
	*goavro.Codec
}

func NewAvroEncoder(schema string) (AvroEncoder, error) {
	var encoder, err = goavro.NewCodec(schema)
	if err != nil {
		return AvroEncoder{}, err
	}

	return AvroEncoder{encoder}, nil
}

func (ae AvroEncoder) Encode(buf []byte, data any) ([]byte, error) {
	return ae.BinaryFromNative(buf, data)
}

type AvroDecoder struct {
	*goavro.Codec
}

func NewAvroDecoder(schema string) (AvroDecoder, error) {
	var decoder, err = goavro.NewCodec(schema)
	if err != nil {
		return AvroDecoder{}, err
	}

	return AvroDecoder{decoder}, nil
}

func (ad AvroDecoder) Decode(buf []byte) (interface{}, error) {
	var data, _, err = ad.NativeFromBinary(buf)
	return data, err
}
