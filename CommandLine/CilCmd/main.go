package  main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

type Config struct {
	username string
	password string
}

var config Config
func main(){
	app := cli.NewApp()
	app.Name = "App helper"
	app.Usage = "Help"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"username,user,u",
			Usage: "User account",
		},
		cli.StringFlag{
			Name:"password,p",
			Usage : "user password",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error{
	config = Config{
		username: c.String("username"),
		password: c.String("password"),
	}

	return exec()
}

func exec() error {
	fmt.Println("username",config.username)
	fmt.Println("password",config.password)

	return nil
}