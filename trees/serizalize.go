package trees

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func Serialize(tree *Node, writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(tree); err != nil {
		return err
	}
	return nil
}

func ToFile(tree *Node, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return Serialize(tree, f)
}

func Deserialize(reader io.Reader) (*Node, error) {
	tree := Tree(0, nil, nil)
	if err := json.NewDecoder(reader).Decode(tree); err != nil {
		return nil, err
	}
	tree.ForEach(func(n *Node) {
		n.Init()
	})
	return tree, nil
}

func FromFile(filename string) (*Node, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Deserialize(bytes.NewReader(file))
}