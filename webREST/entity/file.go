package entity

type File struct {
	FileName    string `json:"filename"`
	Description string `json:"description"`
	FilePath    string `json:"filepath"`
}
