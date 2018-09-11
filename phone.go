package isbr

import (
	"github.com/go-ozzo/ozzo-validation"
)

var (
	Phone = validation.NewStringRule(isBrPhone, "must be a valid phone")
	CellPhone = validation.NewStringRule(isBrCellPhone, "must be a valid cellphone")
)

func isPhone(value string) bool {
	rePhone := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{4}-[0-9]{4}$`)
	return rePhone.MatchString(value)
}

func isCellPhone(value string) bool {
	reCellPhone := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{5}-[0-9]{4}$`)
	return reCellPhone.MatchString(value)
}