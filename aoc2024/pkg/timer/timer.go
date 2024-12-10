package timer

import (
	"fmt"
	"time"
)

func Time(name string) func() {
	start := time.Now()

	return func() {
		elapsed := time.Since(start)

		fmt.Printf("%s took %s\n", name, elapsed)
	}
}
