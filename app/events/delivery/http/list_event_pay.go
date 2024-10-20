package http

import (
	"net/http"
	"strconv"

	"github.com/hammer-code/lms-be/domain"
	"github.com/hammer-code/lms-be/utils"
	"github.com/sirupsen/logrus"
)

func (h Handler) ListEventPay(w http.ResponseWriter, r *http.Request) {
	flterPagination, err := domain.GetPaginationFromCtx(r)
	if err != nil {
		logrus.Error("failed to get pagination : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	startDate, _ := utils.ParseDate(r.URL.Query().Get("start_date"))
	endDate, _ := utils.ParseDate(r.URL.Query().Get("end_date"))

	eventIDs := r.URL.Query().Get("event_id")

	var eventID uint
	if eventIDs != "" {
		eventIDU, err := strconv.ParseUint(eventIDs, 10, 32)
		if err != nil {
			logrus.Error("failed to convert string to uint: ", err)
			utils.Response(domain.HttpResponse{
				Code:    500,
				Message: err.Error(),
			}, w)
			return
		}

		eventID = uint(eventIDU)
	}

	data, pagination, err := h.usecase.ListEventPay(r.Context(), domain.EventFilter{
		ID:               eventID,
		Status:           r.URL.Query().Get("status"),
		StartDate:        startDate,
		EndDate:          endDate,
		FilterPagination: flterPagination,
	})

	if err != nil {
		logrus.Error("failed to list event : ", err)
		utils.Response(domain.HttpResponse{
			Code:    500,
			Message: err.Error(),
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:       200,
		Message:    "success",
		Data:       data,
		Pagination: pagination,
	}, w)
}
