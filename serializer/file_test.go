package serializer_test

import (
	"klever-api/sample"
	"klever-api/serializer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFilename := "../tmp/exchange.bin"

	exchange := sample.NewExchange()

	err := serializer.WriteProtoToBinFile(exchange, binaryFilename)

	require.NoError(t, err)
}