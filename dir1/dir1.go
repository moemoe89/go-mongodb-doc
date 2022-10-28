package dir1

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IntegerCollection = "integers"
)

type Integer struct {
	ID                     primitive.ObjectID     `bson:"_id"`
	SignedInteger          SignedInteger          `bson:"signed_integer"`
	SignedPointerInteger   SignedPointerInteger   `bson:"signed_pointer_integer"`
	UnSignedInteger        UnSignedInteger        `bson:"unsigned_integer"`
	UnSignedPointerInteger UnSignedPointerInteger `bson:"unsigned_pointer_integer"`
}

type SignedInteger struct {
	Int   int   `bson:"int"`
	Int8  int8  `bson:"int8"`
	Int16 int16 `bson:"int16"`
	Int32 int32 `bson:"int32"`
	Int64 int64 `bson:"int64"`
}

type SignedPointerInteger struct {
	IntPointer   *int   `bson:"int_pointer"`
	Int8Pointer  *int8  `bson:"int8_pointer"`
	Int16Pointer *int16 `bson:"int16_pointer"`
	Int32Pointer *int32 `bson:"int32_pointer"`
	Int64Pointer *int64 `bson:"int64_pointer"`
}

type UnSignedInteger struct {
	Uintptr uintptr `bson:"uintptr"`
	Uint    uint    `bson:"uint"`
	Uint8   uint8   `bson:"uint8"`
	Uint16  uint16  `bson:"uint16"`
	Uint32  uint32  `bson:"uint32"`
	Uint64  uint64  `bson:"uint64"`
}

type UnSignedPointerInteger struct {
	Uintptr       *uintptr `bson:"uintptr"`
	UintPointer   *uint    `bson:"uint_pointer"`
	Uint8Pointer  *uint8   `bson:"uint8_pointer"`
	Uint16Pointer *uint16  `bson:"uint16_pointer"`
	Uint32Pointer *uint32  `bson:"uint32_pointer"`
	Uint64Pointer *uint64  `bson:"uint64_pointer"`
}
