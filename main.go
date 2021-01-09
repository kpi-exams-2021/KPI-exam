package main

import (
	"tree/cmd/server"
)

func main() {
	//tree := trees.Tree(10,
	//	trees.Tree(20,
	//		trees.Tree(10, nil, nil),
	//		nil),
	//	trees.Tree(30, nil, nil),
	//)
	//
	//err := serialization.ToFile(tree, "tree.json")
	//if err != nil {
	//	panic(err)
	//}
	//
	//tree, err = serialization.FromFile("tree.json")
	//if err != nil {
	//	panic(err)
	//}
	//
	//tree.Sum()
	//err = serialization.ToFile(tree, "tree2.json")
	//if err != nil {
	//	panic(err)
	//}

	server.RunServer()
}
