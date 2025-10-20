package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// implementing this method means when it's time to convert to json, go will use the return value from calling this method
// as the json representation of whatever the value is.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// we need to quote the string when it's being returned.
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
