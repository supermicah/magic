package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/supermicah/magic/cmd/test"
)

func GetList() []*cli.Command {
	return []*cli.Command{
		test.Test(),
	}
}
