package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" parameter from the current request context, then convert it
// to an integer and return it. If the operation isn't successful, return 0 and
// an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// When httprouter is parsing a request, any interpolated URL parameters
	// will be stored in the request context. We can use the ParamsFromContext()
	// function to retrieve a slice containing these parameter names and values.
	params := httprouter.ParamsFromContext(r.Context())

	// We can then use the ByName() method to get the value of the "id"
	// parameter from the slice. In our project all movies will have a unique
	// positive integer ID, but the value returned by ByName() is always a
	// string. So we try to convert it to a base 10 integer (with a bit size of
	// 64). If the parameter couldn't be converted, or is less than 1, we know
	// the ID is invalid so we return 0 and a new error.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define an envelope type
type envelope map[string]interface{}

// Change the data parameter to have the type envelope instead of interface{}
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
