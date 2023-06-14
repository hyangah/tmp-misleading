//go:build tag1

package mod1

import "fmt"

func Name() string { return fmt.Sprintf("hello %v", "world") }
