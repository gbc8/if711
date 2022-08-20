package Servidor

import (
	"fmt"
	"net"
)

func IniciarServidorUDP() {

	s, err := net.ResolveUDPAddr("udp4", "localhost:5433")
	fmt.Println("Porta UDP 5433")
	if err != nil {
		fmt.Println(err)
		return
	}

	con, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer con.Close()
	buffer := make([]byte, 1024)
	for {
		n, addr, err := con.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		text := string(buffer[0 : n-1])
		go func(text string, addr *net.UDPAddr) {
			data := []byte(RemoverVogais(text))
			_, err = con.WriteToUDP(data, addr)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(text, addr)
	}

}
