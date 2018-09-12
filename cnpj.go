package isbr

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

// CNPJ validate if a string is a correct CNPJ format or not
var CNPJ = validation.NewStringRule(isCNPJ, "must be a valid CNPJ")

func isCNPJ(value string) bool {
	reCNPJ := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}\/\d{4}\-\d{2}$`)
	return reCNPJ.MatchString(value)
}
