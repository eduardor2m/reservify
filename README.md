# Hotel Room Reservation API - Reservify

**Desenvolvido por:** Eduardo Melo (https://github.com/eduardor2m)

Uma API de reserva de quartos para hotéis com autenticação e diferentes níveis de usuário. Este projeto é um exemplo prático de como criar uma API robusta em Go (Golang) usando autenticação JWT e armazenamento de dados no PostgreSQL.

## Recursos

- Autenticação segura usando tokens JWT (JSON Web Tokens).
- Diferentes níveis de usuário: Administrador e Cliente.
- Listagem e detalhes de quartos disponíveis.
- Reserva, atualização e cancelamento de reservas.
- Gerenciamento de quartos e reservas para administradores.

## Pré-requisitos

- Go (Golang) instalado: [Download](https://golang.org/dl/)
- Docker instalado: [Download](https://www.docker.com/)

## Instalação

1. Clone este repositório:

```bash
git clone https://github.com/eduardor2m/reservify.git
```

2. Acesse o diretório do projeto:

```bash
cd reservify
```

3. Instale as dependências:

```bash
go get ./...
```

4. Configure as variáveis de ambiente:

Renomei o arquivo `.env-example` para `.env` no diretório `./cmd/application/` do projeto e configure as seguintes variáveis:

```env
DB_HOST=seu_host_do_banco_de_dados
DB_PORT=5432
DB_USER=seu_usuario_do_banco_de_dados
DB_PASSWORD=sua_senha_do_banco_de_dados
DB_NAME=nome_do_banco_de_dados
JWT_SECRET=seu_segredo_jwt
```

## Uso

1. Execute a aplicação:

```bash
go run ./cmd/application/main.go
```

2. Acesse a API em `http://localhost:8080/api/docs/index.html`.

3. Consulte a documentação interativa da API para obter detalhes sobre os endpoints e como usá-los.

## Documentação

A documentação da API foi gerada automaticamente e está disponível em `http://localhost:8080/api/docs/index.html` após iniciar a aplicação.

## Contribuição

Sinta-se à vontade para contribuir com melhorias, correções de bugs ou novos recursos. Basta seguir as diretrizes de contribuição neste [arquivo](CONTRIBUTING.md).

## Licença

Este projeto está licenciado sob a [Licença XYZ](LICENSE).
```

Lembre-se de substituir `[Seu Nome]`, `[SEU_USUARIO]`, `[nome_do_banco_de_dados]`, `[seu_segredo_jwt]` e outras informações relevantes com os detalhes do seu projeto. Certifique-se de incluir um arquivo de licença e um arquivo `CONTRIBUTING.md` para orientações sobre como os outros podem contribuir para o seu projeto.
