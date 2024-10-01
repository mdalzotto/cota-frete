# API de Cotação de Frete

Este projeto é uma API para cotação de frete, construída utilizando o framework **Gin** no **Golang**, com **MongoDB** para armazenar as cotações.

## Sumário

- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Iniciando o Projeto](#iniciando-o-projeto)
- [Rotas da API](#rotas-da-api)

## Requisitos

Certifique-se de que você tem os seguintes itens instalados no seu ambiente:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://golang.org/dl/) (opcional, caso queira rodar localmente sem Docker)
- [MongoDB](https://www.mongodb.com/try/download/community) (opcional, caso não use o container MongoDB)

## Instalação

Clone o repositório:

   ```bash
   git clone https://github.com/mdalzotto/cota-frete.git
   cd cota-frete
   ```

## Configuração

Crie um arquivo **.env** na raiz do projeto para configurar as variáveis de ambiente necessárias. Use o exemplo abaixo ou o arquivo **env.example** detro do projeto como base:

   ```bash
    API_PORT=8080
    MONGO_HOST=mongodb
    MONGO_PORT=27017
    MONGO_DATABASE=cotacoes
    API_PATH=url_da_api_cotacao
    API_TOKEN=sua_api_token
    API_PLATFORM_CODE=seu_codigo_plataforma
    API_REGISTERED_NUMBER=seu_numero_registrado
    API_DISPATCHER_ZIPCODE=seu_cep
   ```

## Iniciando o Projeto

Com o Docker Compose configurado e todas as configurações feitas, execute o seguinte comando para iniciar o projeto:

   ```bash
   docker-compose up --build
   ```

Certifique-se de que ele esteja rodando.

A API será iniciada em http://localhost:8080.


## Rotas da API

 - Também a um arquivo **Swagger** no projeto que pode ser usado como guia das rotas, caso não queira apenas diga o que esta abaixo. 

### **[GET]** **/metrics?last_quotes={?}**

Para testar o endpoint de métricas, você pode fazer uma requisição GET:

- **Descrição**: Retorna as métricas das cotações de frete.

- **Parâmetro opcional**: last_quotes (int) — Quantidade de cotações retornadas, ordenadas por ordem decrescente.

Exemplo de uso:

  ```bash
  curl "http://localhost:<API_PORT>/metrics?last_quotes=2"
  ```
Isso retornará as métricas das últimas 5 cotações. Se o parâmetro last_quotes não for fornecido, todas as cotações serão consideradas.


### **[POST]** **/quote**

Endpoint utilizado para criar uma cotação de frete. O corpo da requisição deve ser enviado no formato JSON como o exemplo abaixo:
Exemplo de uso:

  ```bash
   {
   "recipient": {
      "address": {
         "zipcode": "01311000"
      }
   },
   "volumes": [
      {
         "category": 7,
         "amount": 1,
         "unitary_weight": 5,
         "price": 349,
         "sku": "abc-teste-123",
         "height": 0.2,
         "width": 0.2,
         "length": 0.2
      },
      {
         "category": 7,
         "amount": 2,
         "unitary_weight": 4,
         "price": 13,
         "sku": "abc-teste-527",
         "height": 0.4,
         "width": 0.6,
         "length": 0.15
      }
   ]
}

  ```

Exemplo de requisição usando curl:

  ```bash
   curl -X POST http://localhost:<API_PORT>/quote \
-H 'Content-Type: application/json' \
-d '{
   "recipient": {
      "address": {
         "zipcode": "01311000"
      }
   },
   "volumes": [
      {
         "category": 7,
         "amount": 1,
         "unitary_weight": 5,
         "price": 349,
         "sku": "abc-teste-123",
         "height": 0.2,
         "width": 0.2,
         "length": 0.2
      },
      {
         "category": 7,
         "amount": 2,
         "unitary_weight": 4,
         "price": 556,
         "sku": "abc-teste-527",
         "height": 0.4,
         "width": 0.6,
         "length": 0.15
      }
   ]
}'
  ```