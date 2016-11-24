package main

import (
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
		cfg, err := config.New(c)

		if err != nil {
			return err
		}

		ctx, err := application.Load(cfg)

		if err != nil {
			return err
		}

		if c.String("start") == "run" {
			return server.Run(ctx)
		} else if c.String("start") == "migrate" {
			return server.Migrate(ctx)
		}

		return failures.ErrBadFlagCli
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
		cli.StringFlag{
			Name:  "start, s",
			Value: "run",
			Usage: "choose between run or migrate. Run start the app and migrate build new tables in the database",
		},
	}

	app.Run(os.Args)
}

// TODO : format error -> faillure, BIND, Validate -> manager
