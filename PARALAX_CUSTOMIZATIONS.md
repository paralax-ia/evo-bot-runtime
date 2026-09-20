# Customizações da Paralax IA

Este repositório é um fork do `evo-bot-runtime` (versão 1.1.0+). Abaixo está o registro das funcionalidades e melhorias customizadas que mantemos ativamente em relação à versão *upstream* da comunidade.

## 1. Preservação de Formatação no DispatchEngine (Mensagens Segmentadas)
A comunidade adotou o nosso design para o `DispatchEngine`, que segmenta respostas longas da IA e processa os envios em partes com postbacks HTTP. No entanto, a implementação do *upstream* utiliza a função `strings.Fields()` nativa do Go para quebrar o texto.

**Problema do Upstream:**
A função `strings.Fields()` consome indiscriminadamente todos os espaços em branco, incluindo as quebras de linha (`\n`). Quando a IA escreve uma mensagem segmentada em parágrafos, o código do upstream concatena e destrói as quebras de linha, enviando um texto chapado.

**Nossa customização:**
- No arquivo `pkg/dispatch/service/dispatch_engine.go`, utilizamos uma Expressão Regular robusta (`regexp.MustCompile(\S+|\s+).FindAllString`) para extrair os tokens.
- O loop de remontagem preserva 100% da formatação e espaçamentos nativos criados pelo LLM, permitindo que mensagens complexas, com listas ou parágrafos extensos, sejam entregues formatadas corretamente aos clientes via Evolution API.
