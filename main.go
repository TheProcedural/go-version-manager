package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"gitlab.com/bjerke-tek/gov/constants"
	"gitlab.com/bjerke-tek/gov/controllers"
)

func main() {
	app := &cli.App{
		Name:      constants.Name,
		Usage:     fmt.Sprintf("version: %s", constants.Version),
		UsageText: "gov [command] [value_if_needed]",
		Commands: []*cli.Command{
			{
				Name:    "list-supported",
				Aliases: []string{"s"},
				Usage:   "List available supported Go versions",
				Action: func(c *cli.Context) error {
					controllers.ListSupportedVersions()
					return nil
				},
			},
			{
				Name:    "list-all",
				Aliases: []string{"a"},
				Usage:   "List all available Go versions",
				Action: func(c *cli.Context) error {
					controllers.ListAllVersions()
					return nil
				},
			},
			{},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install a Go version or use the 'latest' tag",
				Action: func(c *cli.Context) error {
					if c.NArg() == 1 {
						controllers.InstallGoVersion(c.Args().First())
					} else {
						fmt.Println("Please provide a Go version to install.")
					}
					return nil
				},
			},
			{
				Name:    "reinstall",
				Aliases: []string{"r"},
				Usage:   "Reinstall a Go version",
				Action: func(c *cli.Context) error {
					if c.NArg() == 1 {
						controllers.ReinstallGoVersion(c.Args().First())
					} else {
						fmt.Println("Please provide a Go version to reinstall.")
					}
					return nil
				},
			},
			{
				Name:    "use",
				Aliases: []string{"u"},
				Usage:   "Use a specific Go version",
				Action: func(c *cli.Context) error {
					if c.NArg() == 1 {
						controllers.UseGoVersion(c.Args().First())
					} else {
						fmt.Println("Please provide a Go version to use.")
					}
					return nil
				},
			},
			{},
			{
				Name:    "current",
				Aliases: []string{"c"},
				Usage:   "Display the currently used Go version",
				Action: func(c *cli.Context) error {
					controllers.DisplayCurrentVersion()
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List installed Go versions",
				Action: func(c *cli.Context) error {
					controllers.ListInstalledVersions()
					return nil
				},
			},
			{},
			{
				Name:    "remove",
				Aliases: []string{"x"},
				Usage:   "Remove a Go version",
				Action: func(c *cli.Context) error {
					if c.NArg() == 1 {
						controllers.RemoveGoVersion(c.Args().First())
					} else {
						fmt.Println("Please provide a Go version to remove.")
					}
					return nil
				},
			},
			{
				Name:    "prune",
				Aliases: []string{"p"},
				Usage:   "Remove all Go versions except the currently used one",
				Action: func(c *cli.Context) error {
					controllers.Prune()
					return nil
				},
			},
			{},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show the installed version of gov",
				Action: func(c *cli.Context) error {
					controllers.GovVersion()
					return nil
				},
			},
			{
				Name:    "self-update",
				Aliases: []string{"e"},
				Usage:   "Update gov to the latest version",
				Action: func(c *cli.Context) error {
					controllers.SelfUpdate()
					return nil
				},
			},
			{
				Name:  "sayonara",
				Usage: "Uninstall gov from your system",
				Action: func(c *cli.Context) error {
					controllers.Sayonara()
					return nil
				},
			},
			{},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Requieres rework
	// controllers.CheckForGovUpdates()
	// controllers.CheckForGoUpdates()
}
