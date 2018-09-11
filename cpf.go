package isbr

import (
	"github.com/go-ozzo/ozzo-validation"
)

var CPF = validation.NewStringRule(isCPF, "must be a valid CPF")

func isCPF(value string) bool {
	reCPF := regexp.MustCompile(`^([0-9]{3}\.){2}[0-9]{3}-[0-9]{2}$`)
	return reCPF.MatchString(value)
}