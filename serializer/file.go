package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToBinFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("cannot Marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func WriteBinToProto(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data to proto message: %w", err)
	}

	return nil
}
