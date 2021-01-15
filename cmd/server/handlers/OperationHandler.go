package handlers

import (
	"net/http"
	"tree/ttrees"
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
	tree, err := ttrees.FromFile(filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	tree.Op()
	err = ttrees.Serialize(tree, rw)
	if err != nil {
		utils.WriteError(rw, err)
	}
}

func handleOpPost(rw http.ResponseWriter, r *http.Request, filename string) {
	tree, err := ttrees.Deserialize(r.Body)
	if err != nil {
		utils.WriteError(rw, err)
	}

	tree.Op()
	err = ttrees.ToFile(tree, filename)
	if err != nil {
		utils.WriteError(rw, err)
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
}
