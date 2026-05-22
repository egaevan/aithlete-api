package handler

import (
	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/aithlete/aithlete-api/pkg/mock"
	"github.com/labstack/echo/v4"
)

type GoalHandler struct {
	provider *mock.MockProvider
}

func NewGoalHandler(provider *mock.MockProvider) *GoalHandler {
	return &GoalHandler{provider: provider}
}

func (h *GoalHandler) GetGoals(c echo.Context) error {
	data := h.provider.GetGoals()
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *GoalHandler) GetGoal(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.GetGoal(id)
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *GoalHandler) CreateGoal(c echo.Context) error {
	var req request.CreateGoalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.CreateGoal(req.Title, req.Type, req.Target, req.Unit, req.Period, req.Deadline)
	return response.Success(c, 201, code.Success, message.MsgCreated, data)
}

func (h *GoalHandler) UpdateGoal(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateGoalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.UpdateGoal(id, req.Title, req.Type, req.Target, req.Current, req.Unit, req.Period, req.Deadline)
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *GoalHandler) DeleteGoal(c echo.Context) error {
	data := h.provider.DeleteGoal()
	return response.Success(c, 200, code.Success, message.MsgDeleted, data)
}

func (h *GoalHandler) ToggleGoal(c echo.Context) error {
	id := c.Param("id")
	data := h.provider.ToggleGoal(id)
	return response.SuccessOK(c, 200, code.Success, data)
}

func (h *GoalHandler) UpdateGoalProgress(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateGoalProgressRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	data := h.provider.UpdateGoalProgress(id, req.Current)
	return response.SuccessOK(c, 200, code.Success, data)
}
