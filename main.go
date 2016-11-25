package main

import (
	"context"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/juliengroch/todolist/application"
	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/failures"
	"github.com/juliengroch/todolist/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "todolis"
	app.Usage = "formation project"
	app.Action = func(c *cli.Context) error {
		return failures.ErrWrongStartCmdCli
	}

	cf := []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Start the app",
			Flags:   cf,
			Action: func(c *cli.Context) error {
				ctx, err := loadAllContext(c)

				if err != nil {
					return err
				}

				return server.Run(ctx)
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Build new tables in the database",
			Flags:   cf,
			Action: func(c *cli.Context) error {
				ctx, err := loadAllContext(c)

				if err != nil {
					return err
				}

				return application.Migrate(ctx)
			},
		},
	}

	app.Run(os.Args)
}

func loadAllContext(c *cli.Context) (context.Context, error) {
	cfg, err := config.New(c)

	if err != nil {
		return nil, err
	}

	ctx, err := application.Load(cfg)

	if err != nil {
		return nil, err
	}

	return ctx, err
}
