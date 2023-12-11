# GoCmdUtil

## Descrição
Este é um utilitário de linha de comando em Go que permite executar diversas tarefas com a passagem de parâmetros.

## Instalação
Certifique-se de ter o Go instalado em seu sistema. Para instalar as dependências, execute:

```bash
go mod tidy
````

## Execução
### Exemplo 1 - Servidores
#### Para listar servidores, execute:

```bash
go run main.go servidores --host devbook.com.br
```

## Exemplo 2 - Informações de IP
### Para obter informações de IP, execute:

```bash
go run main.go ip --host www.google.com
```

### Comandos Disponíveis
#### servidores
- --host (Obrigatório): Especifica o host para listar os servidores.
#### ip
- --host (Obrigatório): Especifica o host para obter informações de IP.
### Contribuição
### Sinta-se à vontade para contribuir ou reportar problemas. Abra uma issue ou envie um pull request!

## Licença
#### Este projeto é licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.
