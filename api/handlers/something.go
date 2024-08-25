package handlers

import (
	"fmt"
	"log"
	"strconv"

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

func (h *Handler) GetSomethingList(c *gin.Context) {
	// limit and offset
	var req models.SomethingListRequest
	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	req.Limit = limit
	req.Offset = offset

	req.Offset, err = h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	log.Println(req)

	resp, err := h.strg.SomethingI().GetAllSomething(c.Request.Context(), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h *Handler) GetSomethingByID(c *gin.Context) {
	var req models.SomethingReadRequest
	id := c.Param("id")

	req.SomethingId, _ = strconv.Atoi(id)

	fmt.Println(req)

	log.Println(req)

	resp, err := h.strg.SomethingI().GetByIDSomething(c.Request.Context(), &req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h *Handler) DeleteSomething(c *gin.Context) {
	var something models.SomethingDelete
	err := c.ShouldBindJSON(&something)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.strg.SomethingI().DeleteSomething(c.Request.Context(), &something)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h *Handler) UpdateSomething(c *gin.Context) {
	var something models.SomethingUpdate
	err := c.ShouldBindJSON(&something)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	fmt.Println(&something)

	resp, err := h.strg.SomethingI().UpdateSomething(c.Request.Context(), &something)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h *Handler) PatchSomething(c *gin.Context) {
	var something models.SomethingPatch
	err := c.ShouldBindJSON(&something)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.strg.SomethingI().PatchSomething(c.Request.Context(), &something)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
