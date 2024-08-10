package zk

import (
	"fmt"
)

func PanicToError(r any, stack []byte) error {
	return fmt.Errorf("error is %v stack is %s", r, string(stack))
}
