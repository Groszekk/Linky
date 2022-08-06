package controllers

import (
	"Linky/ShorterService/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *LinkController) ShortLink(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}

	short := models.Short{}
	err = json.Unmarshal(body, &short)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}

	success, endpoint := c.Short(short)
	if !success {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	fmt.Fprintf(w, "http://localhost/s/%s", endpoint)
}
