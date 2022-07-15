package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Home !!"))
	if err != nil {
		return
	}
}
