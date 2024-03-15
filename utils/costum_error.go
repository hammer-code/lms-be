package utils

import (
	"strings"

	"github.com/hammer-code/lms-be/domain"
)

func CheckError(err, sub, message string, code int) (domain.HttpResponse, bool) {

	if strings.Contains(err, sub) {
		return domain.HttpResponse{
			Code:    code,
			Message: message,
		}, true
	}
	return domain.HttpResponse{}, false
}

func CostumErr(err string) domain.HttpResponse {

	resp, ok := CheckError(err, "duplicate", "User already exist", 400)
	if ok {
		return resp
	}

	resp, ok = CheckError(err, "password", "Sorry, your password is incorrect", 400)
	if ok {
		return resp
	}

	return domain.HttpResponse{
		Code:    400,
		Message: "Internal server error",
	}

}
