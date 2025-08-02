# Nome da sua imagem Docker
IMAGE_NAME := golang-api-study

# Nome do seu contêiner
CONTAINER_NAME := golang-api-instance

# Porta do contêiner para mapeamento
PORT := 8080

# Compila e constrói a imagem Docker
build:
	docker build -t $(IMAGE_NAME) .

# Executa o contêiner a partir da imagem
run:
	docker run -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

# Remove o contêiner
stop:
	docker stop $(CONTAINER_NAME)

# Limpa o projeto (remove contêiner e imagem)
clean:
	docker rm -f $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)

# Inicia o contêiner em modo interativo (útil para debug)
sh:
	docker run -it --entrypoint /bin/sh $(IMAGE_NAME)


# Teste o endpoint de saúde da API
health:
	curl http://localhost:$(PORT)/health