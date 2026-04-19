package modules

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"tunnelrat/internal/core"
)

// Brute agora testa se o site é VULNERÁVEL a ataques de força bruta
func Brute(targetURL string) core.Result {
	// 1. Configura um cliente com timeout curto para detectar lentidão (throttling)
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	const testLimit = 10
	var responseTimes []time.Duration
	blocked := false

	// Teste de estresse rápido para verificar ausência de Rate Limiting
	for i := 0; i < testLimit; i++ {
		start := time.Now()

		// Enviamos dados genéricos de login
		data := url.Values{}
		data.Set("user", fmt.Sprintf("test_user_%d", i))
		data.Set("pass", "test_pass_123")

		req, _ := http.NewRequest("POST", targetURL, strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "TunnelRat/1.1")

		resp, err := client.Do(req)
		if err != nil {
			// Se o servidor fechar a conexão, ele pode ter um firewall ativo
			blocked = true
			break
		}
		
		responseTimes = append(responseTimes, time.Since(start))
		
		// Se recebermos 429 (Too Many Requests) ou 403 (Forbidden), há proteção
		if resp.StatusCode == 429 || resp.StatusCode == 403 {
			blocked = true
			resp.Body.Close()
			break
		}
		resp.Body.Close()
	}

	// 2. Análise Profissional dos Resultados
	if blocked {
		return core.Result{
			Module: "Brute Force",
			Found:  false,
			Detail: "Proteção detectada (WAF/Rate Limit ativo).",
		}
	}

	// Se todas as requisições responderam com tempo similar e sem erro, é vulnerável
	averageTime := calculateAverage(responseTimes)
	
	return core.Result{
		Module: "Brute Force",
		Found:  true,
		Detail: fmt.Sprintf("VULNERÁVEL: Sem bloqueio após %d tentativas (Média: %v)", testLimit, averageTime),
	}
}

func calculateAverage(times []time.Duration) time.Duration {
	if len(times) == 0 {
		return 0
	}
	var total time.Duration
	for _, t := range times {
		total += t
	}
	return total / time.Duration(len(times))
}