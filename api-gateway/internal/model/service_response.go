package model

type ServiceResponse struct {
	ResponseSuccess bool        `json:"response_success"`
	ResponseMessage string      `json:"response_message"`
	ResponseData    interface{} `json:"response_data"`
}
