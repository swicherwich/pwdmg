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
				Usage:   "TODO usage",
				Subcommands: []*cli.Command{
					{
						Name:    "pwd",
						Aliases: []string{"d"},
						Usage:   "Get password by domain and login",
						Action: func(c *cli.Context) error {
							domain := c.Args().Get(0)
							login := c.Args().Get(1)

							if domain == "" || login == "" {
								return errors.New("usage: pwdmg get pwd <domain> <login>")
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
				Name:    "save",
				Aliases: []string{"s"},
				Usage:   "Save password for provided domain and login account",
				Subcommands: []*cli.Command{
					{
						Name:        "pwd",
						Usage:       "pwdmg save <domain> <login>",
						Description: "Save password for provided domain and login account",
						Action: func(c *cli.Context) error {
							domain := c.Args().Get(0)
							login := c.Args().Get(1)

							if domain == "" || login == "" {
								fmt.Println("Usage: pwdmg save <domain> <login>")
								return nil
							}

							fmt.Print("Password: ")
							pwdB, err := terminal.ReadPassword(syscall.Stdin)
							pwd := string(pwdB)

							if err != nil {
								fmt.Println("Error reading password: ", err)
								os.Exit(1)
							}

							if pwd == "" {
								fmt.Println("Invalid pwd")
								os.Exit(1)
							}

							save.PersistAccount(domain, login, pwd)
							return nil
						},
					},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "TODO usage",
				Subcommands: []*cli.Command{
					{
						Name:        "domain",
						Aliases:     []string{"d"},
						Usage:       "Usage: pwdmg list domain",
						Description: "",
						Action:      func(c *cli.Context) error { return nil },
					},
					{
						Name:        "account",
						Aliases:     []string{"acc"},
						Usage:       "Usage: pwdmg list account",
						Description: "",
						Action:      func(c *cli.Context) error { return nil },
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
