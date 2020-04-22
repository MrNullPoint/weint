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
		Version: "0.0.1",
		Name:    "A simple tool to get somebody's weibo data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "user",
				Aliases:  []string{"u"},
				Usage:    "set weibo user id, must set",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "set to get user's profile",
			},
			&cli.BoolFlag{
				Name:    "weibo",
				Aliases: []string{"w"},
				Usage:   "set to get user's weibo list",
			},
			&cli.BoolFlag{
				Name:    "quick",
				Aliases: []string{"q"},
				Usage:   "set to use quick mode, best practice is to use a proxy pool when set this flag",
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
				Name:        "elastic",
				Aliases:     []string{"e"},
				Usage:       "set elastic search address",
				Value:       "127.0.0.1:9200",
				DefaultText: "127.0.0.1:9200",
			},
		},
	}

	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		spider := weint.NewSpider()
		spider.Uid(c.String("user"))

		if c.Bool("info") {
			spider.Type(weint.TYPE_INFO)
		}

		if c.Bool("weibo") {
			spider.Type(weint.TYPE_WEIBO)
		}

		if c.Bool("quick") {
			spider.Quick(true)
		}

		var filename string
		if c.String("file") == "" {
			filename = "output." + c.String("out")
		} else {
			filename = c.String("file")
		}

		switch c.String("out") {
		case "":
		case "csv":
			spider.Out(&weint.FileCSVOut{FileOut: weint.FileOut{Filename: filename}})
		case "json":
			spider.Out(&weint.FileJsonOut{FileOut: weint.FileOut{Filename: filename}})
		case "db":
			spider.Out(&weint.SQLiteOut{DBName: filename})
		case "elastic":
			spider.Out(&weint.ElasticOut{Host: c.String("elastic")})
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
