package model

type ServiceRequest struct {
	ServiceName   string      `json:"service_name"`
	ServiceMethod string      `json:"service_method"`
	ServiceData   interface{} `json:"service_data"`
}
