package main

import (
	"fmt"
	"github.com/swicherwich/pwdmg/cmd/pwdmg/config"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "pwdmg",
		Usage: "Manage passwords",
		Commands: []*cli.Command{
			config.GetCommand(),
			config.SaveCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
