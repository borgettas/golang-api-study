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
	docker network create go-mysql-network || true
	
	docker run -d -p 3306:3306 \
	--name mysql-instance \
	--network go-mysql-network \
	-e MYSQL_ROOT_PASSWORD=sua_senha_do_mysql \
	-e MYSQL_DATABASE=seu_banco \
	mysql:8.0
	
	docker run -d -p 8080:8080 \
	--name go-api-instance \
	--network go-mysql-network \
	-e DB_USER=root \
	-e DB_PASSWORD=sua_senha_do_mysql \
	-e DB_HOST=mysql-instance \
	-e DB_PORT=3306 \
	-e DB_NAME=seu_banco \
	$(IMAGE_NAME)

# 	docker run -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

# 	docker run -it --rm -p $(PORT):$(PORT) --name $(CONTAINER_NAME) \
# 	-e DB_USER=seu_usuario \
# 	-e DB_PASSWORD=sua_senha_do_usuario \
# 	-e DB_HOST=localhost \
# 	-e DB_PORT=3306 \
# 	-e DB_NAME=seu_banco \
# 	$(IMAGE_NAME)
# 	--entrypoint /bin/sh 
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