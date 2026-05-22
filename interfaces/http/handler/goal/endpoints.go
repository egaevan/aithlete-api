package goal

import (
	"errors"
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/request"
	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) List(c echo.Context) error {
	userID := c.Get("user_id").(string)

	results, err := h.listGoalsUseCase.List(c.Request().Context(), userID)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponseList(results))
}

func (h *Handler) Get(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	result, err := h.getGoalUseCase.Get(c.Request().Context(), userID, goalID)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponse(result))
}

func (h *Handler) Create(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var req request.CreateGoalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.createGoalUseCase.Create(c.Request().Context(), userID, req.Title, req.Type, req.Target, req.Unit, req.Period, req.Deadline)
	if err != nil {
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusCreated, code.Success, message.MsgCreated, toGoalResponse(result))
}

func (h *Handler) Update(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	var req request.UpdateGoalRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.updateGoalUseCase.Update(c.Request().Context(), userID, goalID, req.Title, req.Type, req.Target, req.Current, req.Unit, req.Period, req.Deadline)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponse(result))
}

func (h *Handler) Delete(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	err := h.deleteGoalUseCase.Delete(c.Request().Context(), userID, goalID)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.Success(c, http.StatusOK, code.Success, message.MsgDeleted, nil)
}

func (h *Handler) Toggle(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	result, err := h.toggleGoalUseCase.Toggle(c.Request().Context(), userID, goalID)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponse(result))
}

func (h *Handler) UpdateProgress(c echo.Context) error {
	userID := c.Get("user_id").(string)
	goalID := c.Param("id")

	var req request.UpdateGoalProgressRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, message.MsgBadRequest)
	}

	result, err := h.updateGoalProgressUseCase.UpdateProgress(c.Request().Context(), userID, goalID, req.Current)
	if err != nil {
		if errors.Is(err, domainerr.ErrGoalNotFound) {
			return response.NotFound(c, message.MsgNotFound)
		}
		return response.InternalServerError(c, message.MsgInternalError)
	}

	return response.SuccessOK(c, http.StatusOK, code.Success, toGoalResponse(result))
}
