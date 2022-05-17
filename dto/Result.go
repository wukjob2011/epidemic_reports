package dto

type UpResult struct {
	Success      bool   `json:"success"`
	Data         string `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	RecordCount  int    `json:"RecordCount"`
}
