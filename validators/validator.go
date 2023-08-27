package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors() {

}

func Mobile(f validator.FieldLevel) bool {
	mobile := f.Field().String()

	ok, _ := regexp.MatchString(`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[0-35-9]\d{2}|4(?:0\d|1[0-2]|9\d))|9[0-35-9]\d{2}|6[2567]\d{2}|4[579]\d{2})\d{6}$`, mobile)

	if !ok {
		return false
	}

	return true
}
