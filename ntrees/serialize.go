package ntrees

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func Serialize(tree *NNode, writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(tree); err != nil {
		return err
	}
	return nil
}

func ToFile(tree *NNode, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return Serialize(tree, f)
}

func Deserialize(reader io.Reader) (*NNode, error) {
	tree := NTree(0, nil)
	if err := json.NewDecoder(reader).Decode(tree); err != nil {
		return nil, err
	}
	return tree, nil
}

func FromFile(filename string) (*NNode, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Deserialize(bytes.NewReader(file))
}
