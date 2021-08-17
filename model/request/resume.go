package request

type Resume struct {
	ResumeName string `json:"resume_name" form:"resume_name"`
}
