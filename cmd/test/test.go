package test

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/supermicah/magic/pkg/stdinx"
)

func Test() *cli.Command {
	return &cli.Command{
		Name: "test",
		Action: func(c *cli.Context) error {
			// 假设这是你的原始数据
			data := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}

			// 打印原始数据
			tips := "原始数据：\n"
			for _, item := range data {
				tips = fmt.Sprintf("%s%s\n", tips, item)
			}
			value, err := stdinx.ReadStr(c.Context, tips)
			if err != nil {
				return err
			}
			fmt.Println(value)
			return nil
		},
	}
}
