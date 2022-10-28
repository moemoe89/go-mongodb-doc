## MongoDB Collections

<details>
  <summary>Expand the entities</summary>

<!-- start collection doc -->
### Integer
#### *collection: integers*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "signed_integer": {
    "int": 100,
    "int16": 100,
    "int32": 100,
    "int64": 100,
    "int8": 100
  },
  "signed_pointer_integer": {
    "int16_pointer": 100,
    "int32_pointer": 100,
    "int64_pointer": 100,
    "int8_pointer": 100,
    "int_pointer": 100
  },
  "unsigned_integer": {
    "uint": 300,
    "uint16": 300,
    "uint32": 300,
    "uint64": 300,
    "uint8": 300,
    "uintptr": 824635954880
  },
  "unsigned_pointer_integer": {
    "uint16_pointer": 300,
    "uint32_pointer": 300,
    "uint64_pointer": 300,
    "uint8_pointer": 300,
    "uint_pointer": 300,
    "uintptr": 824635954880
  }
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|signed_integer|<a href="#SignedInteger">SignedInteger</a>|
|signed_pointer_integer|<a href="#SignedPointerInteger">SignedPointerInteger</a>|
|unsigned_integer|<a href="#UnSignedInteger">UnSignedInteger</a>|
|unsigned_pointer_integer|<a href="#UnSignedPointerInteger">UnSignedPointerInteger</a>|

### SignedInteger
| Field | Type |
| --- | --- |
|int|<a href="https://pkg.go.dev/builtin#int">int</a>|
|int8|<a href="https://pkg.go.dev/builtin#int8">int8</a>|
|int16|<a href="https://pkg.go.dev/builtin#int16">int16</a>|
|int32|<a href="https://pkg.go.dev/builtin#int32">int32</a>|
|int64|<a href="https://pkg.go.dev/builtin#int64">int64</a>|

### SignedPointerInteger
| Field | Type |
| --- | --- |
|int_pointer|<a href="https://pkg.go.dev/builtin#int">*int</a>|
|int8_pointer|<a href="https://pkg.go.dev/builtin#int8">*int8</a>|
|int16_pointer|<a href="https://pkg.go.dev/builtin#int16">*int16</a>|
|int32_pointer|<a href="https://pkg.go.dev/builtin#int32">*int32</a>|
|int64_pointer|<a href="https://pkg.go.dev/builtin#int64">*int64</a>|

### UnSignedInteger
| Field | Type |
| --- | --- |
|uintptr|<a href="https://pkg.go.dev/builtin#uintptr">uintptr</a>|
|uint|<a href="https://pkg.go.dev/builtin#uint">uint</a>|
|uint8|<a href="https://pkg.go.dev/builtin#uint8">uint8</a>|
|uint16|<a href="https://pkg.go.dev/builtin#uint16">uint16</a>|
|uint32|<a href="https://pkg.go.dev/builtin#uint32">uint32</a>|
|uint64|<a href="https://pkg.go.dev/builtin#uint64">uint64</a>|

### UnSignedPointerInteger
| Field | Type |
| --- | --- |
|uintptr|<a href="https://pkg.go.dev/builtin#uintptr">*uintptr</a>|
|uint_pointer|<a href="https://pkg.go.dev/builtin#uint">*uint</a>|
|uint8_pointer|<a href="https://pkg.go.dev/builtin#uint8">*uint8</a>|
|uint16_pointer|<a href="https://pkg.go.dev/builtin#uint16">*uint16</a>|
|uint32_pointer|<a href="https://pkg.go.dev/builtin#uint32">*uint32</a>|
|uint64_pointer|<a href="https://pkg.go.dev/builtin#uint64">*uint64</a>|

### Array
#### *collection: arrays*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "arrs": [
    {
      "id": "lorem ipsum"
    }
  ],
  "arrs_pointer": [
    {
      "id": "lorem ipsum"
    }
  ],
  "bools": [
    true
  ],
  "bools_pointer": [
    true
  ],
  "bytes": "Yg==",
  "bytes_pointer": "Yg==",
  "runes": [
    9796
  ],
  "runes_pointer": [
    9796
  ],
  "strings": [
    "lorem ipsum"
  ],
  "strings_pointer": [
    "lorem ipsum"
  ],
  "times": [
    "2022-10-28T12:17:49.4056+09:00"
  ],
  "times_pointer": [
    "2022-10-28T12:17:49.405604+09:00"
  ]
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|bytes|<a href="https://pkg.go.dev/builtin#byte">[]byte</a>|
|runes|<a href="https://pkg.go.dev/builtin#rune">[]rune</a>|
|strings|<a href="https://pkg.go.dev/builtin#string">[]string</a>|
|bools|<a href="https://pkg.go.dev/builtin#bool">[]bool</a>|
|times|<a href="https://pkg.go.dev/time#Time">[]time.Time</a>|
|bytes_pointer|<a href="https://pkg.go.dev/builtin#byte">[]*byte</a>|
|runes_pointer|<a href="https://pkg.go.dev/builtin#rune">[]*rune</a>|
|strings_pointer|<a href="https://pkg.go.dev/builtin#string">[]*string</a>|
|bools_pointer|<a href="https://pkg.go.dev/builtin#bool">[]*bool</a>|
|times_pointer|<a href="https://pkg.go.dev/time#Time">[]*time.Time</a>|
|arrs|<a href="#Arr">[]Arr</a>|
|arrs_pointer|<a href="#Arr">[]*Arr</a>|

### Arr
| Field | Type |
| --- | --- |
|id|<a href="https://pkg.go.dev/builtin#string">string</a>|

### Float
#### *collection: floats*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "float32": 99.99,
  "float32_pointer": 99.99,
  "float64": 99.99,
  "float64_pointer": 99.99
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|float32|<a href="https://pkg.go.dev/builtin#float32">float32</a>|
|float64|<a href="https://pkg.go.dev/builtin#float64">float64</a>|
|float32_pointer|<a href="https://pkg.go.dev/builtin#float32">*float32</a>|
|float64_pointer|<a href="https://pkg.go.dev/builtin#float64">*float64</a>|

### Other
#### *collection: others*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "bool": true,
  "bool_pointer": true,
  "byte": 97,
  "byte_pointer": 97,
  "rune": 9796,
  "rune_pointer": 9796,
  "string": "lorem ipsum",
  "string_pointer": "lorem ipsum"
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|byte|<a href="https://pkg.go.dev/builtin#byte">byte</a>|
|rune|<a href="https://pkg.go.dev/builtin#rune">rune</a>|
|string|<a href="https://pkg.go.dev/builtin#string">string</a>|
|bool|<a href="https://pkg.go.dev/builtin#bool">bool</a>|
|byte_pointer|<a href="https://pkg.go.dev/builtin#byte">*byte</a>|
|rune_pointer|<a href="https://pkg.go.dev/builtin#rune">*rune</a>|
|string_pointer|<a href="https://pkg.go.dev/builtin#string">*string</a>|
|bool_pointer|<a href="https://pkg.go.dev/builtin#bool">*bool</a>|

### Time
#### *collection: times*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "time": "2022-10-28T12:17:49.406301+09:00",
  "time_pointer": "2022-10-28T12:17:49.406301+09:00"
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|time|<a href="https://pkg.go.dev/time#Time">time.Time</a>|
|time_pointer|<a href="https://pkg.go.dev/time#Time">*time.Time</a>|

### Internal
#### *collection: internal*
```json
{
  "_id": "633fe183150872f57b930ac8",
  "integer": {
    "_id": "633fe183150872f57b930ac8",
    "signed_integer": {
      "int": 100,
      "int16": 100,
      "int32": 100,
      "int64": 100,
      "int8": 100
    },
    "signed_pointer_integer": {
      "int16_pointer": 100,
      "int32_pointer": 100,
      "int64_pointer": 100,
      "int8_pointer": 100,
      "int_pointer": 100
    },
    "unsigned_integer": {
      "uint": 300,
      "uint16": 300,
      "uint32": 300,
      "uint64": 300,
      "uint8": 300,
      "uintptr": 824635954560
    },
    "unsigned_pointer_integer": {
      "uint16_pointer": 300,
      "uint32_pointer": 300,
      "uint64_pointer": 300,
      "uint8_pointer": 300,
      "uint_pointer": 300,
      "uintptr": 824635954560
    }
  },
  "time": {
    "_id": "633fe183150872f57b930ac8",
    "time": "2022-10-28T12:17:49.406315+09:00",
    "time_pointer": "2022-10-28T12:17:49.406315+09:00"
  }
}
```
| Field | Type |
| --- | --- |
|_id|<a href="https://pkg.go.dev/go.mongodb.org/mongo-driver/bson/primitive#ObjectID">primitive.ObjectID</a>|
|time|<a href="#Time">Time</a>|
|integer|<a href="#Integer">Integer</a>|


<!-- end collection doc -->

</details>
