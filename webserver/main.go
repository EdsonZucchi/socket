package main

import (
	"fmt"
	"net"
)

func main() {
	host := "localhost"
	port := "8080"
	addr := host + ":" + port

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor : ", err)
	}
	defer listener.Close()

	fmt.Printf("Servidor escutando em %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão ", err)
		}
		fmt.Println("Nova conexão de %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Erro ao ler os dados: ", err)
		return
	}

	fmt.Printf("Dados recebidos: %s\n", string(buff))

	response := "Pong"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Erro ao enviar a resposta : ", err)
		return
	}
}