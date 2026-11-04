package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/model"
)

type DispatchHandler struct {
}

func NewDispatchHandler() *DispatchHandler {
	return &DispatchHandler{}
}

func (h *DispatchHandler) HandleDispatch(c *gin.Context) {
	var req model.ServiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ServiceResponse{
			ResponseSuccess: false,
			ResponseMessage: "Invalid request payload",
			ResponseData:    nil,
		})
		return
	}

	if req.ServiceName == "" || req.ServiceMethod == "" {
		c.JSON(http.StatusBadRequest, model.ServiceResponse{
			ResponseSuccess: false,
			ResponseMessage: "Invalid request service and method",
			ResponseData:    nil,
		})
		return
	}

	// simulated response success
	c.JSON(http.StatusOK, model.ServiceResponse{
		ResponseSuccess: true,
		ResponseMessage: "Request dispatched successfully",
		ResponseData:    nil,
	})
}
