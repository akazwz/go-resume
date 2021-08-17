package request

type BasicInfo struct {
	Name string `json:"name" form:"name" binding:"required"`
}
