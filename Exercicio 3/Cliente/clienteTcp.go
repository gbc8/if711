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
		texto := "Lorem ipsum dolor sit amet, consectetur adipisci elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur. Quis aute iure reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
		c.Write([]byte(texto + "\n"))

		_, _ = bufio.NewReader(c).ReadString('\n')
		// resp, _ := bufio.NewReader(c).ReadString('\n')
		// fmt.Println("Sem vogal:", resp)
	}
	c.Close()
	tempo := time.Since(comeco)
	fmt.Println("Cliente TCP acabou com: ", tempo)
}
