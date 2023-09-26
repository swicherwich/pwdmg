package config

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/swicherwich/pwdmg/internal/app/command/get"
	"github.com/swicherwich/pwdmg/internal/app/command/importpwd"
	"github.com/swicherwich/pwdmg/internal/app/command/remove"
	"github.com/swicherwich/pwdmg/internal/app/command/save"
	"github.com/swicherwich/pwdmg/internal/app/command/update"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func PwdCommand() *cli.Command {
	return &cli.Command{
		Name:  "pwd",
		Usage: "pwdmg pwd",
		Subcommands: []*cli.Command{
			getCommand(),
			saveCommand(),
			updateCommand(),
			removeCommand(),
			importCommand(),
		},
	}
}

func getCommand() *cli.Command {
	return &cli.Command{
		Name:        "get",
		Aliases:     []string{"g"},
		Usage:       "pwdmg pwd get <domain> <login>",
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
	}
}

func removeCommand() *cli.Command {
	return &cli.Command{
		Name:        "remove",
		Aliases:     []string{"rm"},
		Usage:       "pwdmg pwd remove <domain> <login>",
		Description: "Remove domain and login record",
		Action: func(c *cli.Context) error {
			domain := c.Args().Get(0)
			login := c.Args().Get(1)

			if domain == "" || login == "" {
				return errors.New("domain or login cannot be empty")
			}

			fmt.Println("Login removed")

			if err := remove.RemoveAccount(domain, login); err != nil {
				return err
			}
			return nil
		},
	}
}

func saveCommand() *cli.Command {
	return &cli.Command{
		Name:        "save",
		Aliases:     []string{"s"},
		Usage:       "pwdmg pwd save <domain> <login>",
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

func updateCommand() *cli.Command {
	return &cli.Command{
		Name:        "update",
		Aliases:     []string{"upd"},
		Usage:       "pwdmg pwd update <domain> <login>",
		Description: "Update password for provided domain and login account",
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
				return errors.New("empty pwd")
			}

			if err := update.UpdatePassword(domain, login, pwd); err != nil {
				return err
			}
			return nil
		},
	}
}

func importCommand() *cli.Command {
	return &cli.Command{
		Name:        "import",
		Aliases:     []string{"i"},
		Usage:       "pwdmg pwd import",
		Description: "Save password for provided domain and login account",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "provider",
				Aliases: []string{"p"},
				Usage:   "passwords import provider (supports only chrome for now)",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "passwords import file (supports only csv format for now)",
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
