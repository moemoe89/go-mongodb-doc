package subdir2

import (
	"github.com/moemoe89/go-mongodb-doc/dir1"
	"github.com/moemoe89/go-mongodb-doc/dir2"
)

const InternalCollection = "internal"

type Internal struct {
	ID      primitive.ObjectID `bson:"_id"`
	Time    *dir2.Time         `bson:"time"`
	Integer dir1.Integer       `bson:"integer"`
}
