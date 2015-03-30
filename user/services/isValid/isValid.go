package isValid

import (
	"github.com/AitorGuerrero/BadassCity/user"
)

func Service(id string) (interface{}, error) {
	result := false
	if user.ExistsUserWithId(id) {
		result = true
	}
	return result, nil
}
