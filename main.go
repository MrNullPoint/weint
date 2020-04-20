package main

import (
	"Weint/weint"
	"errors"
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
				Name:        "info",
				Aliases:     []string{"i"},
				DefaultText: "true",
				Usage:       "set this flag means you can get weibo user profile",
			},
			&cli.BoolFlag{
				Name:        "weibo",
				Aliases:     []string{"w"},
				DefaultText: "false",
				Usage:       "set this flag means you can get user weibo list",
			},
			&cli.StringFlag{
				Name:    "proxy",
				Aliases: []string{"p"},
				Usage:   "set proxy",
			},
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Usage:   "set output type, csv/json/db/elastic",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "set output filename",
			},
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"h"},
				Usage:       "set elastic search address",
				DefaultText: "127.0.0.1:9200",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		spider := weint.NewSpider()
		spider.Uid(c.String("user"))

		if c.Bool("info") {
			spider.Type(weint.TYPE_INFO)
		}

		if c.Bool("weibo") {
			spider.Type(weint.TYPE_WEIBO)
		}

		var filename string
		if c.String("file") == "" {
			filename = ""
		} else {
			filename = c.String("file")
		}

		var out weint.OutInterface

		switch c.String("out") {
		case "csv":
		case "json":
		case "db":
		case "elastic":
			spider.Out(weint.ElasticOut{Host: c.String("host")})
		default:
			return errors.New("output type you input is not supported yet")
		}

		return spider.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
