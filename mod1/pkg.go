package mod1

import (
	"fmt"

	"golang.org/x/mod/module"
)

func F() {
	fmt.Println("A")
	_ = module.Version{}
}
