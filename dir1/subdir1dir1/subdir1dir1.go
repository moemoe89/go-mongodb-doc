package subdir1dir1

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	FloatCollection = "floats"
)

type Float struct {
	ID             primitive.ObjectID `bson:"_id"`
	Float32        float32            `bson:"float32"`
	Float64        float64            `bson:"float64"`
	Float32Pointer *float32           `bson:"float32_pointer"`
	Float64Pointer *float64           `bson:"float64_pointer"`
}
