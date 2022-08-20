package Servidor

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func IniciarServidorTCP() {
	l, err := net.Listen("tcp", "localhost:5434")
	fmt.Println("Porta TCP 5432")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go aceitarCliente(c)
	}

}

func aceitarCliente(conn net.Conn) {
	fmt.Println("Cliente conectado")
	contador := 0
	for {
		requi, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err, "Cliente desconectou! Requisicoes:", contador)
			return
		}
		requi = strings.TrimSpace(requi)
		respString := RemoverVogais(requi)
		contador++
		conn.Write([]byte(respString + "\n"))
	}
}
