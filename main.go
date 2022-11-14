package main

import (
	_ "go_ticket/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	_ "go_ticket/internal/logic"

	"go_ticket/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
