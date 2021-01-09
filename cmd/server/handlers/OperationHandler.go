package handlers

import (
	"net/http"
	"tree/serialization"
	"tree/utils"
)

type OperationHandler http.HandlerFunc

func OperationHttpHandler(filename string) OperationHandler {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleOpGet(rw, filename)
		} else if r.Method == "POST" {
			handleOpPost(rw, r, filename)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleOpGet(rw http.ResponseWriter, filename string) {
	tree, err := serialization.FromFile(filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	tree.Sum()
	err = serialization.Serialize(tree, rw)
	if err != nil {
		utils.WriteError(rw, err)
	}
}

func handleOpPost(rw http.ResponseWriter, r *http.Request, filename string) {
	tree, err := serialization.Deserialize(r.Body)
	if err != nil {
		utils.WriteError(rw, err)
	}

	tree.Sum()
	err = serialization.ToFile(tree, filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
}
