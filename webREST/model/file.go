package model

type Owner struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Key       int8   `json:"key" binding:"gte=1,lte=100"`
	Email     string `json:"email" binding:"required,email"`
}

type File struct {
	FileName    string `json:"filename" binding:"min=2,max=10"`
	Description string `json:"description"`
	FilePath    string `json:"filepath"`
	Creator     Owner  `json:"creator" binding:"required"`
}
