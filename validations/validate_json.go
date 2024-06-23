package validations

import (
	"net/http"

	custom_errors "github.com/educolog9/helpers/errors"
	languages "github.com/educolog9/helpers/languages"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(w http.ResponseWriter, r *http.Request, obj interface{}) bool {
	var errors []string

	contentLanguage := r.Header.Get("Content-Language")
	if contentLanguage == "" {
		contentLanguage = languages.Spanish.String() // Default to Spanish if not provided
	}

	trans, _ := Uni.GetTranslator(contentLanguage)

	if err := Validate.Struct(obj); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				errors = append(errors, e.Translate(trans))
			}
		} else {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		custom_errors.ReturnBadRequestError(w, errors)
		return true
	}

	return false
}
