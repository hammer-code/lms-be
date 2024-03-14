package utils

import (
	"strings"

	"github.com/hammer-code/lms-be/domain"
)

func CheckError(err, sub, message string) (domain.HttpResponse, bool) {
	if strings.Contains(err, sub) {
		return domain.HttpResponse{
			Code:    400,
			Message: message,
		}, true
	}
	return domain.HttpResponse{}, false
}
