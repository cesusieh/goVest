package app

import (
	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "GoVest"
	app.Usage = "Busca por dados de ações"

	app.Commands = []cli.Command{
		{
			Name: "Search",
			Aliases: []string{
				"s",
			},
			Usage: "Busca pelos dados de um ativo especifíco",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "stock",
					Usage: "O ativo a ser buscado",
					Required: true,
				},
			},
			Action: buscarAcao,
		},

		{
			Name: "ViewKey",
			Aliases: []string{
				"vk",
			},
			Usage:  "Exibe a chave de API cadastrada",
			Action: viewKey,
		},

		{
			Name: "RegisterKey",
			Aliases: []string{
				"rk",
			},
			Usage: "Registra uma chave de API para autenticação",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "key",
					Usage: "A chave de API a ser registrada",
					Required: true,
				},
			},
			Action: registerKey,
		},
	}
	return app
}
