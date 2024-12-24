package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func (t *Case) assertNew() *assert.Assertions {
	return assert.New(&t.testing)
}

func (t *Case) assertEqual(actual, expected interface{}) bool {
	equal := t.assertNew().Equal(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual))
	if !equal {
		err := LogPanic(ErrorHandleEqual(expected, actual))
		if err != nil {
			return false
		}
	}
	return true
}
