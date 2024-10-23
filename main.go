package main

import (
	"Model/funRand"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(funRand.Erlang(4.5, 3))
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
