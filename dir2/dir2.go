package dir2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

const TimeCollection = "times"

type Time struct {
	ID          primitive.ObjectID `bson:"_id"`
	Time        time.Time          `bson:"time"`
	TimePointer *time.Time         `bson:"time_pointer"`
}
