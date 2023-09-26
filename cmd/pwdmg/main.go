package main

import (
	"fmt"
	"github.com/swicherwich/pwdmg/cmd/pwdmg/config"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "pwdmg",
		Usage:       "pwdmg",
		Description: "CLI password manager",
		Commands: []*cli.Command{
			config.PwdCommand(),
		},
	}

	fmt.Println(app.Command("pwd").Command("get").Names())

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
