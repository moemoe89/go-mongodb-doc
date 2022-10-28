package sub1subdir1dir

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ArrayCollection = "arrays"
)

type Array struct {
	ID             primitive.ObjectID `bson:"_id"`
	Bytes          []byte             `bson:"bytes"`
	Runes          []rune             `bson:"runes"`
	Strings        []string           `bson:"strings"`
	Bools          []bool             `bson:"bools"`
	Times          []time.Time        `bson:"times"`
	BytesPointer   []*byte            `bson:"bytes_pointer"`
	RunesPointer   []*rune            `bson:"runes_pointer"`
	StringsPointer []*string          `bson:"strings_pointer"`
	BoolsPointer   []*bool            `bson:"bools_pointer"`
	TimesPointer   []*time.Time       `bson:"times_pointer"`
	Arrs           []Arr              `bson:"arrs"`
	ArrsPointer    []*Arr             `bson:"arrs_pointer"`
}

type Arr struct {
	ID string `bson:"id"`
}
