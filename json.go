package optional

import (
	"encoding/json"
	"reflect"
)

var _ json.Marshaler = Optional[any]{}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.valid {
		return json.Marshal(o.value)
	}

	if reflect.TypeOf(o.value).Kind() == reflect.Struct {
		// As of go 1.20, the standard json encoder does not support omitempty for structs and there is no way to
		// override its definition of "empty".
		// see json.isEmptyValue
		// see https://github.com/golang/go/issues/11939#issuecomment-1275681106
		return []byte("{}"), nil
	}

	var zeroValue T
	return json.Marshal(zeroValue)
}

var _ json.Unmarshaler = &Optional[any]{}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.valid = true
	return json.Unmarshal(data, &o.value)
}
