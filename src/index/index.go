package index

import (
	"fmt"
	"net/http"
)

type IndexCtrl struct {
}

func NewCtrl() IndexCtrl {
	return IndexCtrl{}
}

func (i IndexCtrl) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a REST API :-)")
}
