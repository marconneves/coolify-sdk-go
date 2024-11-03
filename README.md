# Coolify SDK for Go

Este é o SDK oficial para a plataforma [Coolify](https://coolify.io/) em Go. Ele permite que você interaja com a API do Coolify para gerenciar equipes, servidores e habilitar ou desabilitar a API diretamente em Go. Este SDK foi projetado para ser usado com o [Terraform Provider Coolify](https://github.com/marconneves/coolify-sdk-go).

## Documentação

- Documentação oficial do Coolify: [Coolify - Doc](https://coolify.io/docs)

## Instalação

Para instalar o SDK em seu projeto Go, você pode seguir as instruções abaixo:

```bash
go get github.com/marconneves/coolify-sdk-go
go mod tidy
```

## Estrutura de Diretórios

- `client.go`: Arquivo principal de inicialização do cliente.
- `api.go`: Implementa os métodos para habilitar e desabilitar a API.
- `team.go`: Implementa a estrutura e métodos para manipular equipes e membros.
- `server.go`: Implementa a estrutura e métodos para manipular servidores (exemplo simplificado).
- `helpers.go`: Funções auxiliares de resposta e de solicitação HTTP.

## Como Usar

### Configurando o Cliente

Para começar a usar o SDK, você precisa criar uma nova instância do cliente `Client`. 

```go
package main

import (
    "fmt"
    "github.com/marconneves/coolify-sdk-go"
)

func main() {
    client := coolify_sdk.NewClient("https://api.coolify.io", "your-api-token")
    
    // Exemplo de uso do API
    status, err := client.Api().Enable()
    if err != nil {
        fmt.Println("Erro ao habilitar a API:", err)
        return
    }
    fmt.Println("Status da API:", *status)
}
```

### Usando o Módulo API

O módulo `Api` permite habilitar e desabilitar a API do Coolify.

```go
// Habilitar a API
status, err := client.Api().Enable()
if err != nil {
    fmt.Println("Erro ao habilitar a API:", err)
} else {
    fmt.Println("Status da API:", *status) // "success" ou "failure"
}

// Desabilitar a API
status, err = client.Api().Disable()
if err != nil {
    fmt.Println("Erro ao desabilitar a API:", err)
} else {
    fmt.Println("Status da API:", *status) // "success" ou "failure"
}
```

### Usando o Módulo de Equipes

Com o módulo de `TeamInstance`, você pode listar equipes, obter detalhes de uma equipe específica e listar membros de uma equipe.

#### Listar Equipes

```go
teams, err := client.Team().List()
if err != nil {
    fmt.Println("Erro ao listar equipes:", err)
    return
}
for _, team := range *teams {
    fmt.Printf("Equipe: %s, ID: %d\n", team.Name, team.Id)
}
```

#### Obter uma Equipe pelo ID

```go
team, err := client.Team().Get(1)
if err != nil {
    fmt.Println("Erro ao obter equipe:", err)
} else {
    fmt.Printf("Equipe: %s, Descrição: %s\n", team.Name, *team.Description)
}
```

#### Listar Membros de uma Equipe

```go
members, err := client.Team().Members(1)
if err != nil {
    fmt.Println("Erro ao listar membros da equipe:", err)
} else {
    for _, member := range *members {
        fmt.Printf("Membro: %s, Email: %s\n", member.Name, member.Email)
    }
}
```

## Exemplos

Aqui estão alguns exemplos de uso do SDK:

### Exemplo Completo

```go
package main

import (
    "fmt"
    "github.com/marconneves/coolify-sdk-go"
)

func main() {
    client := coolify_sdk.NewClient("https://api.coolify.io", "your-api-token")
    
    // Habilitar a API
    status, err := client.Api().Enable()
    if err != nil {
        fmt.Println("Erro ao habilitar a API:", err)
        return
    }
    fmt.Println("Status da API:", *status)

    // Listar Equipes
    teams, err := client.Team().List()
    if err != nil {
        fmt.Println("Erro ao listar equipes:", err)
        return
    }
    for _, team := range *teams {
        fmt.Printf("Equipe: %s, ID: %d\n", team.Name, team.Id)
    }

    // Obter uma Equipe pelo ID
    team, err := client.Team().Get(1)
    if err != nil {
        fmt.Println("Erro ao obter equipe:", err)
    } else {
        fmt.Printf("Equipe: %s, Descrição: %s\n", team.Name, *team.Description)
    }

    // Listar Membros de uma Equipe
    members, err := client.Team().Members(1)
    if err != nil {
        fmt.Println("Erro ao listar membros da equipe:", err)
    } else {
        for _, member := range *members {
            fmt.Printf("Membro: %s, Email: %s\n", member.Name, member.Email)
        }
    }
}
```

## Configuração do Cliente HTTP

Este SDK utiliza o pacote `net/http` do Go para fazer solicitações HTTP à API do Coolify. Ao criar uma instância de `Client`, você precisa fornecer:

- `hostname`: URL da API do Coolify.
- `apiToken`: Token de autenticação para acesso à API.

O método `httpRequest` é utilizado para fazer as requisições HTTP internamente. Abaixo está um exemplo do método `httpRequest` que pode ser implementado no SDK:

```go
func (c *Client) httpRequest(endpoint, method string, body interface{}) (*http.Response, error) {
    url := fmt.Sprintf("%s/%s", c.hostname, endpoint)
    var buf bytes.Buffer
    if body != nil {
        json.NewEncoder(&buf).Encode(body)
    }
    req, err := http.NewRequest(method, url, &buf)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+c.apiToken)
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("API error: %s", resp.Status)
    }
    return resp, nil
}
```

## Requisitos

- Go 1.16 ou superior.

## Contribuição

Sinta-se à vontade para contribuir com este SDK enviando pull requests ou relatando problemas no repositório oficial no GitHub.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

### Links

- Repositório no GitHub: [coolify-sdk-go](https://github.com/marconneves/coolify-sdk-go)
- Documentação do Coolify: [Coolify - Doc](https://coolify.io/docs)

Esse README fornece instruções claras sobre instalação, uso básico, exemplos e links para mais informações.