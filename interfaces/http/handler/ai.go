package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type AIHandler struct {
	provider *mock.MockProvider
}

func NewAIHandler(provider *mock.MockProvider) *AIHandler {
	return &AIHandler{provider: provider}
}

func (h *AIHandler) GetRecommendations(c echo.Context) error {
	data, meta := h.provider.GetAIRecommendations()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *AIHandler) CreateChatSession(c echo.Context) error {
	data := h.provider.CreateChatSession()
	return response.Success(c, 201, code.Success, message.MsgSuccess, data)
}

func (h *AIHandler) GetChatHistory(c echo.Context) error {
	sessionID := c.Param("sessionId")
	data, meta := h.provider.GetChatHistory(sessionID)
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *AIHandler) SendChatMessage(c echo.Context) error {
	sessionID := c.Param("sessionId")
	var req request.CreateChatRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data, meta := h.provider.SendChatMessage(sessionID, req.Message)
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}

func (h *AIHandler) GetFatigueAnalysis(c echo.Context) error {
	data := h.provider.GetFatigueAnalysis()
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *AIHandler) GetRecoveryScore(c echo.Context) error {
	data := h.provider.GetRecoveryScore()
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *AIHandler) GetPlateauDetection(c echo.Context) error {
	data, meta := h.provider.GetPlateauDetection()
	rMeta := &response.Meta{Total: meta.Total, Page: meta.Page, Limit: meta.Limit, TotalPages: meta.TotalPages}
	return response.SuccessWithMeta(c, 200, code.Success, message.MsgSuccess, data, rMeta)
}
