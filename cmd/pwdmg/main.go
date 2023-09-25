package main

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"pwd-manager/internal/app/command/get"
	"pwd-manager/internal/app/command/save"
	"syscall"
)

func main() {
	app := &cli.App{
		Name:  "pwdmg",
		Usage: "Manage passwords",
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "pwdmg get",
				Subcommands: []*cli.Command{
					{
						Name:        "pwd",
						Usage:       "pwdmg get pwd <domain> <login>",
						Description: "Get password by domain and login",
						Action: func(c *cli.Context) error {
							domain := c.Args().Get(0)
							login := c.Args().Get(1)

							fmt.Println(domain, login)

							if domain == "" || login == "" {
								return errors.New("domain or login cannot be empty")
							}

							pwd, err := get.PwdByLogin(domain, login)
							if err != nil {
								return err
							}

							err = clipboard.WriteAll(pwd)
							if err != nil {
								return err
							}

							fmt.Println("Password copied to clipboard")
							return nil
						},
					},
				},
			},
			{
				Name:        "save",
				Aliases:     []string{"s"},
				Usage:       "pwdmg save <domain> <login>",
				Description: "Save password for provided domain and login account",
				Action: func(c *cli.Context) error {
					domain := c.Args().Get(0)
					login := c.Args().Get(1)

					if domain == "" || login == "" {
						return errors.New("domain or login cannot be empty")
					}

					fmt.Print("Password: ")
					pwdB, err := terminal.ReadPassword(syscall.Stdin)
					pwd := string(pwdB)

					if err != nil {
						return errors.New("error reading password")
					}

					if pwd == "" {
						return errors.New("invalid pwd")
					}

					save.PersistAccount(domain, login, pwd)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
