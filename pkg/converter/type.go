package converter

import "encoding/json"

func StructConverter[T any](data any) (*T, error) {
	var result T
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
