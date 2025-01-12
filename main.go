package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/supermicah/magic/cmd"
)

var VERSION = "v0.0.1"

func main() {
	// 刷新日志，当程序退出时，把在缓存中的内容刷到日志中
	defer func() {
		_ = zap.S().Sync()
	}()
	logger, err := zap.NewDevelopmentConfig().Build(zap.WithCaller(false))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	app := cli.NewApp()
	app.Name = "magic"
	app.Version = VERSION
	app.Usage = "A command line tool for tools"
	app.Authors = append(app.Authors, &cli.Author{
		Name:  "micah",
		Email: "micah.shi@gmail.com",
	})
	app.Commands = cmd.GetList()

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
