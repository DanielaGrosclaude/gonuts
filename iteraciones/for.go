package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i += 4 {
		if i == 4 {
			continue
		}

		fmt.Println(i)
	}

}
