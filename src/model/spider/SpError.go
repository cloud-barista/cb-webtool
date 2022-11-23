package spider

type SpError struct {
	Message string `json:"message"`
}

//type SpError struct {
//	Code    string        `json:"code"`
//	Message SpErrorDetail `json:"message"`
//}

type SpErrorDetail struct {
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Status    string `json:"status"`
}
