package converter

import (
	"go.mongodb.org/mongo-driver/bson"
)

func InputToBson(input any) (bson.D, error) {
	bytes, err := bson.Marshal(input)
	if err != nil {
		return nil, err
	}

	var result bson.D
	err = bson.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	var finalResult bson.D
	for _, data := range result {
		if data.Value == "" || data.Value == nil || data.Key == "_id" {
			continue
		}
		finalResult = append(finalResult, data)
	}

	return finalResult, nil
}
