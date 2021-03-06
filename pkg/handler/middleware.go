package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) checkAuth(c *gin.Context) {
	authHeader := c.GetHeader(entities.Header)
	if authHeader == "" {
		newErrorResponse(c, http.StatusMethodNotAllowed, "Empty Authorization Header")
		return
	}

	headerSlice := strings.Split(authHeader, " ")

	if len(headerSlice) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization Header")
		return
	}

	userId, err := h.services.ParseToken(headerSlice[1])

	if err != nil {
		newErrorResponse(c, http.StatusMovedPermanently, fmt.Sprintf("%s : \n %s", constants.InvalidAuthError, err.Error()))
		return
	}

	// TODO - Send UserID As Context Value
	//ctx := context.WithValue(context.Background(), "ID", "ID")

	c.Set(entities.UserCtx, userId)
}

func (h *Handler) SessionManager(c *gin.Context) {
	s := sessions.Default(c)
	sID := s.Get(constants.SessionID)

	if sID == nil {
		newErrorResponse(c, http.StatusNotFound, constants.GetSessionIDError)
		c.Abort()
		return
	}

}
