package isbr

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

// CPF validate if a string is a correct CPF format or not
var CPF = validation.NewStringRule(isCPF, "must be a valid CPF")

func isCPF(value string) bool {
	reCPF := regexp.MustCompile(`^([0-9]{3}\.){2}[0-9]{3}-[0-9]{2}$`)
	return reCPF.MatchString(value)
}
