# RatGo 🐀

O **RatGo** é uma ferramenta de automação para testes de segurança web desenvolvida em Go. Ele permite identificar vulnerabilidades comuns de forma rápida e eficiente através de uma interface simples via terminal.

## 🚀 Funcionalidades

O programa opera seguindo um menu numérico para facilitar o uso:

1. **XSS (Cross-Site Scripting):** Testa a reflexão de scripts maliciosos em parâmetros de busca.
2. **SQL Injection:** Verifica se o banco de dados do site expõe erros de sintaxe.
3. **Crawl:** Mapeia e extrai links internos e externos da página alvo.
4. **Brute Force:** Analisa se o site possui proteções contra ataques de força bruta em formulários.
5. **Tech Detection:** Identifica tecnologias, servidores e cookies utilizados pelo alvo.

## 🛠️ Como usar

Após rodar o programa, basta seguir a sequência:
1. Escolha o número do módulo (1 a 5).
2. Insira a URL do alvo (exemplo: `https://site.com.br`).

```bash
go run ./cmd/tunnelrat/
