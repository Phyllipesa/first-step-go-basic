package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nome de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	// É um slice de ações que a aplicação pode executar
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca oo nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

/*
Função que busca os IPs de um host
command: go run main.go ip --host amazon.com.br
*/
func buscarIps(c *cli.Context) {
	// Pegando o valor do host passado como argumento
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

/*
Função que busca os servidores de um host
command: go run main.go servidores --host amazon.com.br
*/
func buscarServidores(c *cli.Context) {
	// Pegando o valor do host passado como argumento
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
