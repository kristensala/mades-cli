package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kristensala/mades/commands"
	"github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name: commands.Edx,
                Usage: "for edx communications",
                Subcommands: renderSubcommands(commands.Edx),
            },
            {
                Name: commands.Ecp,
                Usage: "form ecp communications",
                Subcommands: renderSubcommands(commands.Ecp),
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

func renderSubcommands(commandName string) []*cli.Command {
    mades := commands.Mades{
        Ecp: "",
        Edx: "",
    }

    subcommands := []*cli.Command{
        {
            Name: "send",
            Usage: "Send a message",
            Flags: []cli.Flag{
                &cli.StringFlag{
                    Name: "message",
                    Aliases: []string{"m"},
                    Usage: "Message you want to send",
                    Required: true,
                },
            },
            Action: func(ctx *cli.Context) error {
                mades.SendMessage(commandName, ctx.String("message"))
                return nil
            },
        },
        {
            Name: "receive",
            Usage: "Receive messages",
            Action: func(ctx *cli.Context) error {
                fmt.Println("")
                return nil
            },
        },
        {
            Name: "confirm",
            Usage: "Confirm a messages",
            Flags: []cli.Flag{
                &cli.StringFlag{
                    Name: "messageId",
                    Aliases: []string{"i"},
                    Usage: "Message id you want to confirm",
                    Required: true,
                },
            },
            Action: func(ctx *cli.Context) error {
                fmt.Println("")
                return nil
            },
        },
    }

    return subcommands
}
