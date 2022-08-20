package Cliente

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func IniciarClienteTCP() {
	c, err := net.Dial("tcp", "localhost:5434")
	if err != nil {
		fmt.Println(err)
		return
	}
	comeco := time.Now()
	for i := 0; i < 10000; i++ {
		texto := "Nelson Rosa"
		c.Write([]byte(texto + "\n"))

		_, _ = bufio.NewReader(c).ReadString('\n')
		//resp, _ := bufio.NewReader(c).ReadString('\n')
		//fmt.Println("Sem vogal:", resp)
	}
	c.Close()
	tempo := time.Since(comeco)
	fmt.Println("Cliente TCP acabou com: ", tempo)
}
