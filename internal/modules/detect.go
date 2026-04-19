package modules

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"tunnelrat/internal/core"
)

// Detect analisa cabeçalhos e cookies para identificar a stack tecnológica
func Detect(targetURL string) core.Result {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return core.Result{
			Module: "Tech Detection",
			Found:  false,
			Detail: fmt.Sprintf("Erro ao conectar: %v", err),
		}
	}
	defer resp.Body.Close()

	var techs []string

	// 1. Análise de Cabeçalhos HTTP (Server, X-Powered-By, etc.)
	headers := resp.Header
	
	if s := headers.Get("Server"); s != "" {
		techs = append(techs, "Servidor: "+s)
	}
	if p := headers.Get("X-Powered-By"); p != "" {
		techs = append(techs, "Linguagem/Framework: "+p)
	}
	if s := headers.Get("X-AspNet-Version"); s != "" {
		techs = append(techs, "ASP.NET v"+s)
	}

	// 2. Análise de Cookies (Assinaturas comuns)
	for _, cookie := range resp.Cookies() {
		cName := strings.ToUpper(cookie.Name)
		if strings.Contains(cName, "PHPSESSID") {
			techs = append(techs, "PHP")
		} else if strings.Contains(cName, "JSESSIONID") {
			techs = append(techs, "Java/JSP")
		} else if strings.Contains(cName, "ASPSESSIONID") || strings.Contains(cName, "ASP.NET_SESSIONID") {
			techs = append(techs, "ASP.NET")
		} else if strings.Contains(cName, "CAKEPHP") {
			techs = append(techs, "CakePHP")
		}
	}

	// 3. Verificação de Segurança (HSTS, CSP)
	if headers.Get("Strict-Transport-Security") == "" {
		techs = append(techs, "HSTS Ausente (Risco)")
	}

	if len(techs) > 0 {
		return core.Result{
			Module: "Tech Detection",
			Found:  true,
			Detail: strings.Join(techs, " | "),
		}
	}

	return core.Result{
		Module: "Tech Detection",
		Found:  false,
		Detail: "Nenhuma assinatura tecnológica óbvia encontrada nos cabeçalhos.",
	}
}