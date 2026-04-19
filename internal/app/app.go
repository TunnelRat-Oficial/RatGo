package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tunnelrat/internal/core"
	"tunnelrat/internal/modules"
	"tunnelrat/internal/ui"
)

type App struct {
	Target string
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	ui.Banner()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(`
[1] Crawl
[2] SQLi Test
[3] XSS Test
[4] Brute Force
[5] Tech Detect
[0] Sair
`)

		fmt.Print("Escolha: ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if op == "0" {
			break
		}

		if a.Target == "" {
			fmt.Print("URL alvo: ")
			url, _ := reader.ReadString('\n')
			a.Target = strings.TrimSpace(url)
		}

		switch op {
		case "1":
			core.PrintResult(modules.Crawl(a.Target))
		case "2":
			core.PrintResult(modules.SQLi(a.Target))
		case "3":
			core.PrintResult(modules.XSS(a.Target))
		case "4":
			core.PrintResult(modules.Brute(a.Target))
		case "5":
			core.PrintResult(modules.Detect(a.Target))
		}
	}
}