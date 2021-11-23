package serializer_test

import (
	"klever-api/pb"
	"klever-api/sample"
	"klever-api/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFilename := "../tmp/exchange.bin"
	jsonFilename := "../tmp/exchange.json"

	exchange1 := sample.NewExchange()
	err := serializer.WriteProtoToBinFile(exchange1, binaryFilename)
	require.NoError(t, err)

	exchange2 := &pb.ExchangeRequest{}
	err = serializer.WriteBinToProto(binaryFilename, exchange2)
	require.NoError(t, err)
	require.True(t, proto.Equal(exchange1, exchange2))

	err = serializer.WriteProtoToJSONFile(exchange1, jsonFilename)
	require.NoError(t, err)
}