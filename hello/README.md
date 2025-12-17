# Hello

Projeto de introdução ao Go, demonstrando o uso de módulos externos e locais.

## 📝 Conceitos Abordados

- Importação de pacotes externos (`rsc.io/quote`)
- Uso de módulos locais com `replace` directive
- Estrutura básica de um programa Go
- Gerenciamento de dependências com `go.mod`

## 🔧 Estrutura

    ```
        hello/
        ├── go.mod      # Definição do módulo e dependências
        └── hello.go    # Código principal
    ```

## 🚀 Como Executar

    ```sh
        go run hello.go
    ```

**Saída esperada:**

    ```
        Don't communicate by sharing memory, share memory by communicating.
        Hi, Bruno. Welcome!
    ```

## 📦 Dependências

- `rsc.io/quote` v1.5.2 - Módulo externo para citações sobre Go
- `example.com/greetings` - Módulo local (definido via replace directive)

## 💡 Aprendizados

- Como inicializar um módulo Go com `go mod init`
- Como adicionar dependências externas com `go get`
- Como referenciar módulos locais usando `replace` no `go.mod`
- Importação e uso de funções de outros pacotes
