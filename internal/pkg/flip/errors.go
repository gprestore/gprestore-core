package flip

type FlipError struct {
	Code   string           `json:"code,omitempty"`
	Errors []map[string]any `json:"errors,omitempty"`
}
