package converter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		if data.Value == "" || data.Value == nil {
			data.Value = nil
			continue
		}
		if data.Key == "_id" {
			objectId, ok := data.Value.(primitive.ObjectID)
			if !ok {
				objectId, err = primitive.ObjectIDFromHex(data.Value.(string))
				if err != nil {
					return nil, err
				}
			}
			data.Value = objectId
		}
		finalResult = append(finalResult, data)
	}

	return finalResult, nil
}
