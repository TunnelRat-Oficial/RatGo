package modules

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"tunnelrat/internal/core"
)

// SQLi verifica se o endpoint reflete erros de sintaxe de banco de dados
func SQLi(targetURL string) core.Result {
	// Payloads de teste clássicos para forçar erros de sintaxe
	payloads := []string{"'", "\"", "';", "')", "''"}
	
	// Assinaturas de erro comuns em diversos bancos de dados
	dbErrors := []string{
		"SQL syntax", "mysql_fetch", "PostgreSQL query failed",
		"ORA-01756", "SQLite3::SQLException", "Dynamic SQL Error",
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Como o SQLi geralmente ocorre em parâmetros, testamos o 'id' ou 'q'
	param := "id" 

	for _, p := range payloads {
		// Monta a URL maliciosa
		testURL := fmt.Sprintf("%s?%s=%s", targetURL, param, p)
		
		resp, err := client.Get(testURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		content := strings.ToLower(string(body))

		// Analisa se o corpo da página contém erros conhecidos de DB
		for _, dbErr := range dbErrors {
			if strings.Contains(content, strings.ToLower(dbErr)) {
				return core.Result{
					Module: "SQL Injection",
					Found:  true,
					Detail: fmt.Sprintf("VULNERÁVEL (Error-Based): Erro '%s' detectado usando payload [%s]", dbErr, p),
				}
			}
		}
	}

	return core.Result{
		Module: "SQL Injection",
		Found:  false,
		Detail: "Nenhuma vulnerabilidade óbvia de SQLi detectada via erros.",
	}
}