package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type HTTPBathHandler struct {
	userRepo model.UserRepo
}

func NewHTTPBathHandler(repo model.UserRepo) *HTTPBathHandler {
	return &HTTPBathHandler{
		userRepo: repo,
	}
}

func (h *HTTPBathHandler) HandleHTTPBath(c echo.Context) error {
	id := c.FormValue("id")

	id = strings.ReplaceAll(id, " ", "")
	id = strings.TrimSpace(id)

	user, err := h.userRepo.Find(id)
	if err != nil {
		logrus.Infof("couldnt fetch user(id: %s) info from db: %s", id, err.Error())
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}
