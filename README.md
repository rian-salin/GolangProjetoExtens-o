## API Aprendendo Golang

Projeto de API de usuários em Go com MySQL. A camada de modelo aplica hash `bcrypt` antes de persistir senhas.

### Pré-requisitos
- Go 1.23 (ou defina `GOTOOLCHAIN=go1.23.0` nos comandos `go`)
- MySQL acessível (padrão: `treehousedb` em `127.0.0.1:3306`)
- Variáveis de ambiente configuradas em `config/.env`

### Configurando ambiente
1. Copie `config/.env` (já existe um exemplo) e ajuste:
   ```
   API_PORT=:8081
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_ADDR=host:porta
   DB_DATABASE=treehousedb
   SECRET_KEY=alguma_chave
   ```
2. Garanta que o banco tem as tabelas do `sql/ddl_local.sql`.

### Executando
```bash
GOTOOLCHAIN=go1.23.0 go run main.go
```
O log mostrará `Servidor ouvindo em :8081 ...`.

### Testando criação de usuário via curl
Com o servidor ativo, envie:
```bash
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{
        "nome_usuario": "Maria Souza",
        "cpf": "12345678909",
        "email_usuario": "maria@example.com",
        "senha": "minhaSenhaTop123"
      }'
```
O campo `senha` retornará em hash (`$2a$...`). Verifique o banco (`SELECT id,name,password FROM users;`) para confirmar.

### Rodando testes
Não há testes automatizados compilados, mas você pode validar hashing manualmente:
```bash
GOTOOLCHAIN=go1.23.0 go test ./...
```
Implemente testes em `models/users_test.go` para garantir que `security.Hash` e `security.Verify` funcionem como esperado.
