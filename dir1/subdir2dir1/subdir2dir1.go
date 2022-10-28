package subdir2dir1

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	OtherCollection = "others"
)

type Other struct {
	ID            primitive.ObjectID `bson:"_id"`
	Byte          byte               `bson:"byte"`
	Rune          rune               `bson:"rune"`
	String        string             `bson:"string"`
	Bool          bool               `bson:"bool"`
	BytePointer   *byte              `bson:"byte_pointer"`
	RunePointer   *rune              `bson:"rune_pointer"`
	StringPointer *string            `bson:"string_pointer"`
	BoolPointer   *bool              `bson:"bool_pointer"`
}
