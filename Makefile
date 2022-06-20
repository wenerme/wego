REPO_ROOT ?= $(shell git rev-parse --show-toplevel)
-include $(REPO_ROOT)/mod.mk

start-dev-pg:
	docker run -it --rm -p 5432:5432 \
	-e POSTGRES_PASSWORD=dev \
	-e POSTGRES_DB=dev \
	-v $(PWD)/ignored/pg:/var/lib/postgresql/data \
	--name pg postgres:alpine
