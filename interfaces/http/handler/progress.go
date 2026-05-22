package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type ProgressHandler struct {
	provider *mock.MockProvider
}

func NewProgressHandler(provider *mock.MockProvider) *ProgressHandler {
	return &ProgressHandler{provider: provider}
}

func (h *ProgressHandler) GetBodyWeightHistory(c echo.Context) error {
	data, meta := h.provider.GetBodyWeightHistory()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *ProgressHandler) AddBodyWeight(c echo.Context) error {
	var req request.AddBodyWeightRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.AddBodyWeight(req.Weight, req.BodyFatPercentage)
	return response.Success(c, 201, code.Success, message.MsgBodyWeightAdded, data)
}

func (h *ProgressHandler) GetStrengthProgression(c echo.Context) error {
	data, meta := h.provider.GetStrengthProgression()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *ProgressHandler) GetConsistency(c echo.Context) error {
	data, meta := h.provider.GetConsistency()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *ProgressHandler) GetMuscleVolume(c echo.Context) error {
	data, meta := h.provider.GetMuscleVolume()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *ProgressHandler) GetProgressOverview(c echo.Context) error {
	data := h.provider.GetProgressOverview()
	return response.SuccessOK(c, 200, code.Success, data)
}
