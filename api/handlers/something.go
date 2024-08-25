package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-psql/api/http"
	"github.com/go-psql/models"
)

func (h *Handler) CreateSomething(c *gin.Context) {
	var something models.SomethingCreate
	err := c.ShouldBindJSON(&something)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.strg.SomethingI().CreateSomething(c.Request.Context(), &something)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
