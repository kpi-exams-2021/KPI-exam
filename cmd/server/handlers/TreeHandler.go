package handlers

import (
	"net/http"
	"tree/serialization"
	"tree/utils"
)

type TreeHttpHandler http.HandlerFunc

func TreeHandler(filename string) TreeHttpHandler {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGet(rw, filename)
		} else if r.Method == "POST" {
			handlePost(rw, r, filename)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
 	}
}

func handleGet(rw http.ResponseWriter, filename string) {
	tree, err := serialization.FromFile(filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	err = serialization.Serialize(tree, rw)
	if err != nil {
		utils.WriteError(rw, err)
	}
}

func handlePost(rw http.ResponseWriter, r *http.Request, filename string) {
	tree, err := serialization.Deserialize(r.Body)
	if err != nil {
		utils.WriteError(rw, err)
	}

	err = serialization.ToFile(tree, filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
}
