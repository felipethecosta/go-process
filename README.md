# Sistema em Golang para execução de tarefas em segundo plano

## Descrição

O sistema é uma aplicação de execução de tarefas em segundo plano, tendo sua finalidade de apenas aparecer o icone na barra de tarefas do windows, linux e mac.

## Estrutura do Projeto

O projeto está organizado da forma padrão do projeto golang, onde está sem pastas apenas a lógica para aparecer o icone na barra de tarefas

## Como rodar o projeto

Para rodar o projeto, você precisa ter o golang instalado na sua máquina, após isso você precisa rodar o comando abaixo para instalar as dependências do projeto:

```bash
go mod tidy
```

Após isso você precisa rodar o comando abaixo para rodar o projeto:

Para compilar o projeto, você precisa rodar o comando abaixo:
```bash
go build -o meu_programa
```

Após isso você precisa executar o executavel que foi gerado, e para isso você precisa executar o comando abaixo:

```bash
./meu_programa
```
