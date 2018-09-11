package isbr

import (
	"github.com/go-ozzo/ozzo-validation"
)

var CEP = validation.NewStringRule(isCEP, "must be a valid CEP")

func isCEP(value string) bool {
	reCEP := regexp.MustCompile(`^[0-9]{5}-[0-9]{3}$`)
	return reCEP.MatchString(value)
}