package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator"
)

func MinimalValidation(field validator.FieldLevel) bool {
	param, _ := strconv.Atoi(field.Param())
	if len(field.Field().String()) < param {
		return false
	}

	return true
}

func MaximumValidation(field validator.FieldLevel) bool {
	param, _ := strconv.Atoi(field.Param())
	if len(field.Field().String()) > param {
		return false
	}

	return true
}

func Captcha(field validator.FieldLevel) bool {
	captchaSecret := os.Getenv("CAPTCHA")
	responseKey := field.Field().String()

	url := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s", captchaSecret, responseKey)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(body))

	return true
}
