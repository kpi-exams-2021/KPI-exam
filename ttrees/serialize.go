package ttrees

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func Serialize(tree *TNode, writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(tree); err != nil {
		return err
	}
	return nil
}

func ToFile(tree *TNode, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return Serialize(tree, f)
}

func Deserialize(reader io.Reader) (*TNode, error) {
	tree := TTree(0, [3]*TNode{})
	if err := json.NewDecoder(reader).Decode(tree); err != nil {
		return nil, err
	}
	return tree, nil
}

func FromFile(filename string) (*TNode, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Deserialize(bytes.NewReader(file))
}
