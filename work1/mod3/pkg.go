package mod3

import (
	"fmt"

	"example.com/mod2"
)

func B() {
	fmt.Println("C")
	mod2.A()
}
