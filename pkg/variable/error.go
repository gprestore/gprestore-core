package variable

import "fmt"

var (
	ErrUserFilter         = fmt.Errorf("validator: filter id, username, email, or phone not found")
	ErrStoreFilter        = fmt.Errorf("validator: filter id, slug, or author_id not found")
	ErrItemFilter         = fmt.Errorf("validator: filter id, slug, or store_id not found")
	ErrItemFilterStoreId  = fmt.Errorf("validator: filter store_id not found")
	ErrStockFilter        = fmt.Errorf("validator: filter id, store_id, or item_id not found")
	ErrStockFilterItemId  = fmt.Errorf("validator: filter item_id not found")
	ErrOrderFilter        = fmt.Errorf("validator: filter id, code, store_id, or customer.email not found")
	ErrOrderFilterStoreId = fmt.Errorf("validator: filter store_id not found")
	ErrUnauthorized       = fmt.Errorf("unauthorized")
)
