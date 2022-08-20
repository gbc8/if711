package main

import (
	sv "atv4/Servidor"
	"os"
)

func main() {

	args := os.Args
	tipoServidor := args[1]

	if tipoServidor == "tcp" {
		sv.IniciarServidorTCP()
	} else {
		sv.IniciarServidorUDP()
	}

}
