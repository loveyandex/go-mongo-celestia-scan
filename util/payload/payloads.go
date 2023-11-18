package payload



type ApiResponse struct{
	
	Status bool `json:"status"`
	
	Data interface{} `json:"data,omitempty"`

}
type ApiListResponse struct{
	
	Status bool `json:"status"`
	
	Data interface{} `json:"data,omitempty"`
	TotalCount int64 `json:"total_count,omitempty"`

}

type ErrorResponse struct {
	Status bool `json:"status"`
	Message string 
	Code int 
}
