package goutil

import (
	"fmt"
	"time"
)

func Test() string {
	return fmt.Sprintf("goutil test run %v", time.Now())
}
