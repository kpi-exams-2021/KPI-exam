package client

import (
	"log"
	"net/http"
	"os"
	"tree/trees"
)

type Api string

type TreeClient struct {
	api string
	client *http.Client
	filename string
}

func NewClient(api Api, filename string) *TreeClient {
	return &TreeClient{
		api: string(api),
		client: &http.Client{},
		filename: filename,
	}
}

func (c *TreeClient) GetTree() {
	r, err := c.client.Get(c.api)
	if err != nil {
		log.Fatal(err)
		return
	}

	tree, err := trees.Deserialize(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = trees.ToFile(tree, c.filename); err != nil {
		log.Fatal(err)
	}
}

func (c *TreeClient) PostTree() {
	f, err := os.Open(c.filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	r, err := c.client.Post(c.api, "application/json", f)
	if err != nil {
		log.Fatal(err)
		return
	}
	if r.StatusCode != 200 {
		log.Fatalf("Error: %d", r.StatusCode)
	}
	log.Println("POSTed successfully")
}

func (c *TreeClient) GetOp() {
	r, err := c.client.Get(c.api + "op/")
	if err != nil {
		log.Fatal(err)
		return
	}

	tree, err := trees.Deserialize(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = trees.ToFile(tree, c.filename); err != nil {
		log.Fatal(err)
	}
}

func (c *TreeClient) PostOp() {
	f, err := os.Open(c.filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	r, err := c.client.Post(c.api + "op/", "application/json", f)
	if err != nil {
		log.Fatal(err)
		return
	}
	if r.StatusCode != 200 {
		log.Fatalf("Error: %d", r.StatusCode)
	}
	log.Println("POSTed successfully")
}

func (c *TreeClient) GetTreeAndMakeOp() {
	r, err := c.client.Get(c.api)
	if err != nil {
		log.Fatal(err)
		return
	}

	tree, err := trees.Deserialize(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	tree.Sum()
	if err = trees.ToFile(tree, c.filename); err != nil {
		log.Fatal(err)
	}
}
