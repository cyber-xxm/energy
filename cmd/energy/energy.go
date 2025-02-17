//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package main energy command-line
package main

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/jessevdk/go-flags"
	"os"
	"os/signal"
	"syscall"
)

var commands = map[string]*command.Command{
	"install": internal.CmdInstall,
	"package": internal.CmdPackage,
	"version": internal.CmdVersion,
	"env":     internal.CmdEnv,
	"init":    internal.CmdInit,
	"build":   internal.CmdBuild,
	"bindata": internal.CmdBindata,
	"gen":     internal.CmdGen,
	"upg":     internal.CmdUpgrade,
	"help":    internal.CmdHelp,
	"cli":     internal.CmdCli,
}

func main() {
	//term.GoENERGY()
	termRun()
}

func termRun() {
	wd := tools.CurrentExecuteDir()
	cc := &command.Config{Wd: wd}
	parser := flags.NewParser(cc, flags.HelpFlag|flags.PassDoubleDash)
	if len(os.Args) < 2 {
		parser.WriteHelp(term.TermOut)
		os.Exit(0)
	}

	// Go 构建参数, 它应出现在 cli 有效参数最后一个位置
	// 之后的参数做为 go build [args] 传递进去
	buildArgsIdx := tools.GetBuildArgsFlagIndex()
	var args = os.Args[1:]
	if buildArgsIdx > 1 { // energy build --buildargs 索引一定大于 1
		// 拿到 cli 有效参数
		args = os.Args[1:buildArgsIdx]
	}
	if extraArgs, err := parser.ParseArgs(args); err != nil {
		println(err.Error())
		return
	} else {
		cmd, ok := commands[parser.Active.Name]
		if !ok {
			parser.WriteHelp(term.TermOut)
		}
		// energy [cmd] help
		if len(extraArgs) > 0 {
			name := extraArgs[0]
			if name == "help" {
				term.Section.Println(cmd.UsageLine, "\n", cmd.Long)
				os.Exit(0)
			}
		}
		if cmd.Short != "" {
			term.Section.Println(cmd.Short)
		}
		env.InitDevEnvConfig(wd) //初始化本地配置文件
		signalHandler()
		if err := cmd.Run(cc); err != nil {
			term.Logger.Error(err.Error())
			os.Exit(1)
		}
	}
}

func signalHandler() {
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			sig := <-ctrlC
			println(fmt.Sprintf("\nReceived signal: %v. CTRL+C Force Exit.", sig))
			os.Exit(1)
		}
	}()
}
