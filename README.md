# GOLANG/Postgres

Importação de arquivo csv/txt para base de dados Postgresql utilizando a linguagem GOLANG.

### Instruções do projeto

**Obrigatórios:**

- Criar um serviço em GO que receba um arquivo csv/txt de entrada (Arquivo Anexo)
- Este serviço deve persistir no banco de dados relacional (postgres) todos os dados contidos no arquivo
  Obs: O arquivo não possui um separador muito convencional
- Deve-se fazer o split dos dados em colunas no banco de dados
  Obs: pode ser feito diretamente no serviço em GO ou em sql
- Realizar higienização dos dados após persistência (sem acento, maiusculo, etc)
- Validar os CPFs/CNPJs contidos (validos e não validos numericamente)
- Todo o código deve estar disponível em repositório publico do GIT

**Desejável:**

- Utilização das linguagen GOLANG para o desenvolvimento do serviço
- Utilização do DB Postgres
- Docker Compose , com orientações para executar (arquivo readme)

### Como Utilizar

**Dependências:**

- Necessita a linguagem de programação GO instalada
- brdoc package, pacote de terceiro para validação de CPF/CNPJ
- Postgres Database (Sugerida a utilização da ferramenta PGADMIN para consultas a base de dados)
- GIT instalado

Para utilização do projeto:
Clonar o seguinte repositório - git clone https://github.com/erivaltonoliveira/golang-csv-postgres

Há um package necessário para a execução do projeto, contendo funções para validação de documentos CPF e CNPJ. Necessário verificar antes de rodar a aplicação. Sendo necessário executar o seguinte comando na pasta do projeto:

    go get "github.com/Nhanderu/brdoc"

Após clonar o repositório, ir até o diretório /src do projeto e executar o comando go run main.go

Ao acessar http://localhost:8080 (ou APPLICATIONPATH/view/index.html):

- Será carregado formulario para seleção e envio do arquivo texto, que pode ser tanto no formato CSV , quanto TXT

- Ao confirmar o envio a aplicação irá importados para a da base de dados (tabela CLIENTE), validando as informações de documentos como CPF e CNPJ.

### Estrutura relacional

A estrutura relacional para a aplicação consiste de apenas uma tabela [CUSTOMER], e sua crição será efetuada automaticamente na execução da aplicaçõ.

Estrutura da tabela:

```
    CREATE TABLE IF NOT EXISTS CUSTOMER (
    id serial,
    cpf text,
    cpf_valido bool,
    private int,
    incompleto int,
    data_ultima_compra date,
    ticket_medio numeric(15,2),
    ticket_ultima_compra numeric(15,2),
    cnpj_mais_frequente text,
    cnpj_mais_frequente_valido bool,
    cnpj_ultima_compra text,
    cnpj_ultima_compra_valido bool
    )
```

### Configuração de Banco de Dados

```
    HOST     = "localhost"
    PORT     = 5432
    USER     = "postgres"
    PASSWORD = "root"
    DBNAME   = "dbasego"
```

### Estrutura de pastas do projeto

A estrutura de arquivos segue o padrão MVC:

```
.
└── src
    ├── assets
        └── base_teste.txt
    └── control
        └── control.go
    └── db
        └── db.go
    └── lib
        └── util
        	├── convert.go
        	└── file.go
    └── model
        └── customer.go
    └── tmp
        └── tempfile.txt
    └── view
        └── index.html
        ├── main.go
        └── README.md

```

### Para rodar o projeto com Docker composer

Na raíz temos o arquivo "docker-compose.yml" para utilização com Docker (necessário instalar Docker)

Para execução da aplicação executar "docker-compose up" no diretório raíz:

### Fontes de estudo

- https://golang.org/doc/
- https://docs.docker.com/
- https://www.admfactory.com/how-to-convert-a-string-to-an-int-type-in-golang/
- https://tutorialedge.net/golang/go-file-upload-tutorial/
