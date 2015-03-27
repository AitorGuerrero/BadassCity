package isValid

import (
	"github.com/koding/kite"
)

func Service(r *kite.Request) (interface{}, error) {
	result := false
	id := r.Args.One().MustString()
	if len(id) > 0 {
		result = true
	}
	return result, nil
}
