BIN := "./bin/link_remover_tg_bot"
IMAGE := "link_remover_tg_bot:dev"


#build:
#	docker build -t $(IMAGE) --secret id=TOKEN,env=$TOKEN -f Dockerfile .
build:
	DOCKER_BUILDKIT=1 docker build --secret id=TOKEN,env=TOKEN -t $(IMAGE) -f Dockerfile .

run:
	docker run $(IMAGE)

PHONY: build run