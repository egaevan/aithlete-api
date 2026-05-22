package exercise

import (
	"net/http"

	"github.com/aithlete/aithlete-api/interfaces/http/response"
	"github.com/aithlete/aithlete-api/pkg/code"
	"github.com/aithlete/aithlete-api/pkg/message"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListMuscleGroups(c echo.Context) error {
	groups := h.listMuscleGroupsUseCase.ListMuscleGroups()
	resp := toMuscleGroupResponseList(groups)
	meta := response.NewMeta(len(resp), 1, 20)
	return response.SuccessWithMeta(c, http.StatusOK, code.Success, message.MsgSuccess, resp, meta)
}
