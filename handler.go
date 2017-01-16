package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}

func productlist(w http.ResponseWriter, r *http.Request) {

}
