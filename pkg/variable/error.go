package variable

import "fmt"

var (
	ErrFilter       = fmt.Errorf("validator: filter id, username, email, or phone not found")
	ErrUnauthorized = fmt.Errorf("unauthorized")
)
