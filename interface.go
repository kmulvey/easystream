package easystream

type KVPair struct {
	Key   []byte
	Value []byte
}

type Config struct {
	Brokers []string
	TLSCert []byte
	TLSKey  []byte
	Encoder
	Decoder
}

type Publisher interface {
	PublishMessage(KVPair) error
	PublishMessages([]KVPair) error
}

type Encoder interface {
	Encode(any) error
	SetSchema([]byte) error
	GetSchema() []byte
}

type Decoder interface {
	Decode([]byte) (any, error)
	SetSchema([]byte) error
	GetSchema() []byte
}
