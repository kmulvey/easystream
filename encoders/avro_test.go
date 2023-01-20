package encoders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSchema = `{"type": "record", "name": "LoginEvent", "fields": [{"name": "Username", "type": "string"}]}`

func TestAvroEncodeDecode(t *testing.T) {
	t.Parallel()

	var encoder, err = NewAvroEncoder(testSchema)
	assert.NoError(t, err)

	var testData = map[string]interface{}{
		"Username": "superman",
	}

	encodedAvro, err := encoder.Encode(nil, testData)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x10, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e}, encodedAvro)

	decoder, err := NewAvroDecoder(testSchema)
	assert.NoError(t, err)

	decodedData, err := decoder.Decode(encodedAvro)
	assert.NoError(t, err)
	assert.EqualValues(t, testData, decodedData)
}
