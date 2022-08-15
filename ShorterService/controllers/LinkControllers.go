package controllers

import (
	"Linky/ShorterService/interfaces"
	"Linky/ShorterService/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type LinkController struct {
	interfaces.ILinkService
}

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

	endpoint, success := c.Short(short)
	if !success {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w)
		return
	}
	fmt.Fprintf(w, "http://localhost/s/%s", endpoint)
}

func (c *LinkController) ShortRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	link, success := c.GetRedirectLink(short)
	if !success {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
		return
	}

	http.Redirect(w, r, link, http.StatusFound)
}
