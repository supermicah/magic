package stdinx

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func ReadStr(_ context.Context, tips string) (string, error) {
	// 从标准输入读取用户输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(tips)
	fmt.Print("请输入（以回车结束）： ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return "", err
	}
	return strings.TrimSpace(input), nil
}
