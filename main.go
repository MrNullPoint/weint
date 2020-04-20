package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "user",
				Aliases:  []string{"u"},
				Usage:    "set weibo user id",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "set this flag means you can get weibo user profile",
			},
			&cli.BoolFlag{
				Name:    "weibo",
				Aliases: []string{"w"},
				Usage:   "set this flag means you can get user weibo list",
			},
			&cli.StringFlag{
				Name:    "proxy",
				Aliases: []string{"p"},
				Usage:   "set proxy",
			},
			&cli.StringFlag{
				Name:    "proxy",
				Aliases: []string{"p"},
				Usage:   "set proxy",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
