## Curso Rest Api Golang Alura

# Projeto de Personalidades da Marvel com Docker e PostgreSQL

Este projeto utiliza Docker e PostgreSQL para criar um banco de dados de personalidades da Marvel.

## Conteudo 

- Crie uma API do zero com Go
- Integre sua API Go com um banco de dados sendo executado no Docker
- Aprenda a utilizar o GORM, o ORM mais famoso do Go
- Saiba como criar um middleware e evite código duplicado
- Integre sua API Go com um front-end React

## Requisitos

- Docker
- Docker Compose
- Go

## Instruções de Uso

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto.
3. Antes de iniciar a aplicação, é necessário criar o banco de dados. Para isso, navegue até a pasta `infra` e execute `docker-compose up`.
4. O script de inicialização irá criar a tabela `personalidade` e inserir 10 registros de personalidades da Marvel.

## Acessando o Banco de Dados

Para acessar o banco de dados em execução, use o seguinte comando:

``bash 
docker exec -it <container_id_or_name> /bin/bash

Claro, aqui está um exemplo básico de um arquivo README para o seu projeto:

```markdown
# Projeto de Personalidades da Marvel com Docker e PostgreSQL

Este projeto utiliza Docker e PostgreSQL para criar um banco de dados de personalidades da Marvel.

## Requisitos

- Docker
- Docker Compose

## Instruções de Uso

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto.
3. Execute o comando `docker-compose up` para iniciar o serviço do banco de dados.
4. O script de inicialização irá criar a tabela `personalidade` e inserir 10 registros de personalidades da Marvel.

## Acessando o Banco de Dados

Para acessar o banco de dados em execução, use o seguinte comando:

```bash
docker exec -it <container_id_or_name> /bin/bash
```

Substitua `<container_id_or_name>` pelo ID ou nome do seu contêiner.

## Estrutura da Tabela

A tabela `personalidade` possui os seguintes campos:

- `id`: um identificador único para cada personalidade (chave primária).
- `nome`: o nome da personalidade.
- `profissao`: a profissão ou papel da personalidade.
- `historia`: uma breve descrição ou história da personalidade.

## Subindo aplicacao web
Entrar na pasta frontend-react-personalidades-master e executar os comandos abaixo:
```bash
npm install
npm start
```

## Contribuindo

Contribuições são bem-vindas! Por favor, leia as diretrizes de contribuição antes de enviar um pull request.

## Licença

Este projeto está licenciado sob a licença MIT.