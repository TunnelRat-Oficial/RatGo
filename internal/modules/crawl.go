package modules

import (
	"fmt"
	"io"         // Faltava este import!
	"net/http"
	"net/url"
	"strings"
	"time"
	"tunnelrat/internal/core"

	"golang.org/x/net/html"
)

// Crawl realiza uma varredura de links em uma página alvo
func Crawl(targetURL string) core.Result {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return core.Result{
			Module: "Crawl",
			Found:  false,
			Detail: fmt.Sprintf("Erro ao acessar: %v", err),
		}
	}
	defer resp.Body.Close()

	// Extrai os links do corpo da resposta
	links := extractLinks(resp.Body, targetURL)
	
	count := len(links)
	if count > 0 {
		// Mostra a quantidade e os primeiros links encontrados
		limit := 3
		if count < limit {
			limit = count
		}
		summary := strings.Join(links[:limit], " | ")
		return core.Result{
			Module: "Crawl",
			Found:  true,
			Detail: fmt.Sprintf("[%d links] Ex: %s", count, summary),
		}
	}

	return core.Result{
		Module: "Crawl",
		Found:  false,
		Detail: "Nenhum link encontrado.",
	}
}

// extractLinks agora recebe io.Reader diretamente, que é o correto
func extractLinks(body io.Reader, base string) []string {
	var links []string
	z := html.NewTokenizer(body)
	baseURL, _ := url.Parse(base)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						val := a.Val
						u, err := url.Parse(val)
						if err == nil {
							// Resolve caminhos relativos para URLs completas
							abs := baseURL.ResolveReference(u)
							links = append(links, abs.String())
						}
					}
				}
			}
		}
	}
}