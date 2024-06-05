package domain

type (
	Submit struct {
		Language string `json:"lang"`
		Code     string `json:"code"`
		SubmitAt string `json:"submitAt"`
		Status   string `json:"status"`
	}
)

