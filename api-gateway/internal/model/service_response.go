package model

type ServiceResponse struct {
	responseSuccess bool
	responseMessage string
	responseData    interface{}

	ServiceRequest *ServiceRequest
}
