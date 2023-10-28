.PHONY: provision
provision: down build mod-tidy up

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: down-volumes
down-volumes:
	docker-compose down --remove-orphans --volumes

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: reprovision
reprovision: provision

.PHONY: update
update:
	docker-compose run --no-deps --rm app go get -u ./...

.PHONY: sh
sh:
	docker-compose run --no-deps --rm -it app sh

.PHONY: mod-tidy
mod-tidy:
	docker-compose run --no-deps --rm app go mod tidy
