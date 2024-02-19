package cmd

import (
	"github.com/peifengll/gotools/internal/leet"
	"github.com/urfave/cli/v2"
)

func GetApp() *cli.App {
	app := &cli.App{
		UseShortOptionHandling: true,
		Usage:                  "这是一个工具包，使用命令行提供一些基本功能",
		Commands: []*cli.Command{
			{
				Name:    "leet",
				Aliases: []string{"l"},
				Usage:   "create a go file to practice for leetcode",
				Action:  leet.CreateToday,
			},
		},
	}

	return app
}
