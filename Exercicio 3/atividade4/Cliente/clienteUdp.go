package Cliente

import (
	"fmt"
	"net"
	"time"
)

func IniciarClienteUDP() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:5433")
	if err != nil {
		panic(err)
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	comeco := time.Now()
	for i := 0; i < 10000; i++ {
		data := []byte("Nelson Rosa" + "\n")
		_, err = c.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 512)
		_, _, err := c.ReadFromUDP(buffer)
		//t, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println("Sem vogal: ", string(buffer[0:t]))
	}
	tempo := time.Since(comeco)
	fmt.Println("Cliente UDP acabou com: ", tempo)
}
