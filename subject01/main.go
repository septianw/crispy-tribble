package main

import (
	"fmt"
	"os"
	"plugin"
)

type Greeter interface {
	Greet()
}

func main() {
	// determind plugin to load
	var mod string
	var err error
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	default:
		fmt.Println("Don't speak that language")
		os.Exit(1)
	}

	// 1. load module
	// open *.so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// 2. lookup symbols (An exported function and variable
	// In this case variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// 3. Assert that loaded symbol is of desired type
	// in this case interface type Greeter
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(4)
	}

	// 4. use the module
	greeter.Greet()
	fmt.Println("hello world")
}
