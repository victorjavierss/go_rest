package index

import (
	"net/http"
	"fmt"
)

type Index struct {
}

func (i Index) Get (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a REST API :-)")
}