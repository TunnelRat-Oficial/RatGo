package core

import (
	"fmt"
	"tunnelrat/internal/ui"
)

type Result struct {
	Module string
	Found  bool
	Detail string
}

func PrintResult(r Result) {
	if r.Found {
		fmt.Println(ui.Red + "[VULNERABILIDADE ENCONTRADA] " + ui.Reset + r.Module)
		fmt.Println(ui.Red + "DETALHE: " + r.Detail + ui.Reset)
	} else {
		fmt.Println(ui.Green + "[OK] Sistema não apresenta vulnerabilidade: " + r.Module + ui.Reset)
	}
}