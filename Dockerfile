# Use uma única etapa para simplificar e garantir a execução.
FROM golang:1.24-alpine

# Definindo o diretório de trabalho no contêiner
WORKDIR /app

# Copiando todos os arquivos do seu diretório local para o contêiner
COPY . .

# Baixando as dependências e compilando o aplicativo.
# O binário 'main' será criado no diretório /app.
RUN go mod download
RUN go build -o main .

# Expondo a porta que a aplicação usará
EXPOSE 8080

# Comando para executar o aplicativo.
# O binário 'main' agora está garantido para estar no diretório /app.
CMD ["./main"]