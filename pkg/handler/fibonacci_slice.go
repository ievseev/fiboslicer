package handler

import (
	"github.com/gin-gonic/gin"
	fiboslicer "github.com/ievseev/fibonacci-slicer"
	"net/http"
)

func (h *Handler) getFibonacciSlice(c *gin.Context) {
	var sequenceLimit fiboslicer.SequenceLimit

	if err := c.BindJSON(&sequenceLimit); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fiboSlice, err := h.services.GetFibonacciSlice(sequenceLimit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"fibonacci_slice": fiboSlice,
	})
}
