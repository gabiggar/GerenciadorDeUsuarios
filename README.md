# GerenciadorDeUsuarios

API REST em Go para gerenciamento simples de usuários em memória.

## Visão geral

Projeto de exemplo que expõe rotas para criar, listar e consultar usuários. Os dados são armazenados em memória no processo e não persistem entre execuções.

## Requisitos

- Go 1.25+

## Como executar

No diretório do projeto:

```bash
go run .
```

O servidor é iniciado em `http://localhost:8080`.

## Endpoints

### `POST /api/users`

Cria um novo usuário.

Exemplo de corpo esperado:

```json
{
  "first_name": "",
  "last_name": "",
  "biography": ""
}
```

### `GET /api/users`

Retorna todos os usuários cadastrados.

### `GET /api/users/{id}`

Retorna o usuário correspondente ao `id` informado.

### `PUT /api/users/{id}`

Atualiza os dados do usuário identificado pelo `id`.

Exemplo de corpo esperado:

```json
{
  "first_name": "",
  "last_name": "",
  "biography": ""
}
```

### `DELETE /api/users/{id}`

Remove o usuário correspondente ao `id` informado.

## Estrutura do projeto

- `main.go` - inicializa o servidor HTTP.
- `cmd/api/api.go` - define as rotas e middleware usando `chi`.
- `handlers/user.go` - contém os handlers HTTP.
- `models/user.go` - simula um repositório em memória e fornece operações de busca/inserção.
- `dto/` - estrutura de dados para requisições e respostas.