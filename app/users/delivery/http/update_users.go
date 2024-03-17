package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) UpdateProfileUser(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		utils.Response(domain.HttpResponse{
			Code:    401,
			Message: "Not permission",
		}, w)
		return
	}

	id, err:=jwt.ExtractToken(authorizationHeader, config.GetConfig().JWT_SECRET_KEY)
	if err !=nil{
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}



	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	userUpdate := domain.UserUpdateProfile{}
	if err := json.Unmarshal(bodyBytes, &userUpdate); err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	err = h.usecase.UpdateProfileUser(r.Context(), userUpdate, id)

	if err != nil {
		logrus.Error("userUsecase: failed to update users")
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "Success update",
		Data:    userUpdate,
	}, w)
}
