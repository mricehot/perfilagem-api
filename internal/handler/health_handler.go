package handler

import (
	"fmt"
	"net/http"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")

}
