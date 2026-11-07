package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/client"
	"api-gateway/internal/middleware"
	"api-gateway/internal/model"
	"api-gateway/internal/registry"
	"api-gateway/internal/rpctypes"
)

type DispatchHandler struct{}

func NewDispatchHandler() *DispatchHandler {
	return &DispatchHandler{}
}

func (h *DispatchHandler) HandleDispatch(c *gin.Context) {
	var req model.ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	route, ok := registry.Lookup(req.ServiceName, req.ServiceMethod)
	if !ok {
		fail(c, http.StatusBadRequest, "Unknown service or method")
		return
	}

	// service_data (interface{}) -> typed args
	args := route.NewArgs()
	raw, err := json.Marshal(req.ServiceData)
	if err != nil || json.Unmarshal(raw, args) != nil {
		fail(c, http.StatusBadRequest, "Invalid service_data")
		return
	}

	callerID, callerRole := middleware.Identity(c)

	if route.Role != "" {
		if callerID == 0 {
			fail(c, http.StatusUnauthorized, "Authentication required")
			return
		}
		if callerRole != route.Role {
			fail(c, http.StatusForbidden, "Forbidden")
			return
		}
	}

	if scoped, ok := args.(rpctypes.SellerScoped); ok {
		scoped.SetSellerID(callerID)
	}

	reply := route.NewReply()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := client.Call(ctx, route.Address, route.RPCName, args, reply); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fail(c, http.StatusGatewayTimeout, "Service timed out")
			return
		}
		fail(c, http.StatusBadGateway, "Service unavailable")
		return
	}

	status := http.StatusOK
	if !reply.IsSuccess() {
		status = route.FailStatus
		if status == 0 {
			status = http.StatusBadRequest
		}
	}

	c.JSON(status, model.ServiceResponse{
		ResponseSuccess: reply.IsSuccess(),
		ResponseMessage: reply.GetMessage(),
		ResponseData:    reply,
	})
}

func fail(c *gin.Context, status int, msg string) {
	c.JSON(status, model.ServiceResponse{
		ResponseSuccess: false,
		ResponseMessage: msg,
	})
}
