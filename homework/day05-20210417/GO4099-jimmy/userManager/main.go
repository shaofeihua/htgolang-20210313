package main

import (
	"fmt"
	"strconv"
	"strings"
	"userManager/commands"
	"userManager/common"
	_ "userManager/init"
)

func prompt() {
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("1. 退出")
	commands.Commands.Prompts()
	fmt.Println(strings.Repeat("=", 20))
}

func main() {
	if !common.Auth() {
		return
	}
	fmt.Println("欢迎来到用户管理系统")
	for {
		prompt()

		cmd, _ := strconv.Atoi(common.Input("请输入指令: "))
		if callback := commands.Commands.Get(cmd); callback != nil {
			callback()
		} else if cmd == 1 {
			break
		} else {
			fmt.Printf("您输入的指令有误\n")
		}
	}
}
