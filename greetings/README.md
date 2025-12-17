# Greetings

Módulo local simples que fornece funções de saudação personalizadas.

## 📝 Conceitos Abordados

- Criação de um pacote Go reutilizável
- Exportação de funções (convenção de nomenclatura com letra maiúscula)
- Uso de `fmt.Sprintf` para formatação de strings
- Estrutura de um módulo Go

## 🔧 Estrutura

    ```
        greetings/
        ├── go.mod        # Definição do módulo
        └── greetings.go  # Implementação das funções
    ```

## 📚 API

### `Hello(name string) string`

Retorna uma mensagem de saudação personalizada para o nome fornecido.

**Exemplo:**

    ```go
        message := greetings.Hello("Bruno")
        // Retorna: "Hi, Bruno. Welcome!"
    ```

## 💡 Aprendizados

- Como criar um pacote Go exportável
- Convenção de nomenclatura para funções públicas (PascalCase)
- Como estruturar um módulo local para ser usado por outros projetos
- Documentação de funções com comentários
