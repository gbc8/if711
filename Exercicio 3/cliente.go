package main

import (
	cl "atv3/Cliente"
	"os"
)

func main() {

	args := os.Args
	TipoConexao := args[1]

	if TipoConexao == "tcp" {
		cl.IniciarClienteTCP()
	} else {
		cl.IniciarClienteUDP()
	}

}
