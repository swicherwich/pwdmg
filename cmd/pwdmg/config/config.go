package config

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/swicherwich/pwdmg/internal/app/command/get"
	"github.com/swicherwich/pwdmg/internal/app/command/importpwd"
	"github.com/swicherwich/pwdmg/internal/app/command/save"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func GetCommand() *cli.Command {
	return &cli.Command{
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
	}
}

func SaveCommand() *cli.Command {
	return &cli.Command{
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
	}
}

func ImportCommand() *cli.Command {
	return &cli.Command{
		Name:        "import",
		Usage:       "pwdmg import",
		Description: "Save password for provided domain and login account",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "provider",
				Aliases: []string{"p"},
				Usage:   "passwords import provider (for now support only chrome)",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "passwords import file (for now support only csv format)",
			},
		},
		Action: func(c *cli.Context) error {
			p := c.String("provider")
			f := c.String("file")

			switch p {
			case "chrome", "Chrome":
				importpwd.ImportFromChrome(f)
			default:
				fmt.Printf("Provider %s is unsupported\n", p)
			}

			return nil
		},
	}
}
