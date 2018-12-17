package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "faithdroid cli tool"
	app.Usage = "dev faithdroid app faster"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name: "new",
			Action: func(c *cli.Context) error {
				fmt.Println("create project in", c.Args().First())
				return nil
			},
		},
	}
	e := app.Run(os.Args)
	if e != nil {
		fmt.Println(`app.Run error :`, e)
		return
	}
}
