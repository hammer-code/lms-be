package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	loginInstance := domain.Login{}
	if err := json.Unmarshal(bodyBytes, &loginInstance); err != nil {
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	_, token, err := h.usecase.Login(r.Context(), loginInstance)
	if err != nil {
		resp := utils.CostumErr(err.Error())
		utils.Response(resp, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    200,
		Message: "Login successfully",
		Data:    token,
	}, w)
}
