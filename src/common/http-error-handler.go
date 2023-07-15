package common

import (
	"errors"
	"log"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, err error) {
	log.Printf("%v", err)

	if errors.Is(err, NotFoundError) {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	if errors.Is(err, LogicError) {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if errors.Is(err, ForbiddenError) {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	http.Error(w, http.StatusText(500), 500)
}
