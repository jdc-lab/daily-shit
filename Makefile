GO_CMD=go
DOCKER_COMPOSE_CMD=docker-compose
JWT_SECRET_FOR_DEV=a_SUPER_secret

.PHONY: generate
generate:
	@cd user-service && \
	${GO_CMD} generate
	@cd gateway && \
	${GO_CMD} generate

.PHONY: consul-start
consul-start:
	@${DOCKER_COMPOSE_CMD} up -d consul-service

.PHONY: consul-stop
consul-stop:
	@${DOCKER_COMPOSE_CMD} stop consul-service

.PHONY: user-service
user-service:
	@cd user-service && \
	${GO_CMD} run . -jwt-secret="${JWT_SECRET_FOR_DEV}"

.PHONY: gateway
gateway:
	@cd gateway && \
	${GO_CMD} run .

.PHONY: all-up
all-up:
	@${DOCKER_COMPOSE_CMD} up -d --build

.PHONY: all-down
all-down:
	@${DOCKER_COMPOSE_CMD} down