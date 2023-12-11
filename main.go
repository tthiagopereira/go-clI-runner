package main

import (
	app2 "cli-command/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partita")
	app := app2.Gerar()
	erro := app.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
