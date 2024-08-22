package variable

import "fmt"

var (
	ErrUserFilter   = fmt.Errorf("validator: filter id, username, email, or phone not found")
	ErrStoreFilter  = fmt.Errorf("validator: filter id, slug, or author_id not found")
	ErrUnauthorized = fmt.Errorf("unauthorized")
)
