package main

import (
	"fmt"
	"net"
)

func main() {

	addr := "localhost:8080"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Erro ao conectar servidor: ", err)
	}
	defer conn.Close()

	ping := "ping"
	_, err = conn.Write([]byte(ping))
	if err != nil {
		fmt.Println("Erro ao enviar ping: ", err)
		return
	}
	fmt.Println("Ping enviado para o servidor")

	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		fmt.Println("Erro ao receber resposta: ", err)
		return 
	}
	fmt.Printf("Resposta do servidor: %s\n", string(response))

}
