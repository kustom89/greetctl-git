package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	gr "github.com/kustom89/greetctl/greetctl/internal/logx"
	"github.com/kustom89/greetctl/greetctl/pkg/greet"
)

func main() {

	name := flag.String("name", "Mundo", "Nombre de la persona a saludar")
	flag.Parse()

	msg := greet.Hello(*name)

	fmt.Println("sssssss", msg)
	color.Green("This is a green message %s", *name)
	color.Red("This is a green message %s", *name)
	color.Blue("This is a green message %s", *name)
	color.Yellow("This is a green message %s", *name)

	gr.Info("This is an info message")
}
