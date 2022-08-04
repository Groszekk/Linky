package middlewares

import (
	"Linky/ShorterService/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func Validation(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		short := models.Short{}
		err = json.Unmarshal(body, &short)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w)
			return
		}

		validate := validator.New()
		validate.RegisterValidation("min", MinimalValidation)
		validate.RegisterValidation("max", MaximumValidation)

		err = validate.Struct(short)

		var info []models.ValidateInfo

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				switch err.ActualTag() {
				case "min":
					{
						info = append(info, models.ValidateInfo{Field: err.Field(), Message: fmt.Sprintf("%s is too short", err.Field())})
					}

				case "max":
					{
						info = append(info, models.ValidateInfo{Field: err.Field(), Message: fmt.Sprintf("%s is too long", err.Field())})
					}
				}
			}

			jsonRes, _ := json.Marshal(info)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, string(jsonRes))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
