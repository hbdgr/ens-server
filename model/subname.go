package model

type SubnameInfo struct {
	LabelHash   string  `json:"label_hash"`
	SubnameNode string  `json:"node"`
	Owner       string  `json:"owner"`
	Resolver    string  `json:"resolver"`
	Subname     string  `json:"name"`
	ErrorMsg    *string `json:"error_message,omitempty"`
}

type SubnamesInfo struct {
	Parent   string         `json:"parent_name"`
	Subnames []*SubnameInfo `json:"subnames"`
}
