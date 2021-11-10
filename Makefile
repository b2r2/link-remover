BIN := "./bin/link_remover_tg_bot"
IMAGE := "link_remover_tg_bot:dev"


build:
	docker build -t $(IMAGE) --build-arg TOKEN='455278361:AAFSYcbmNvtshujXKU8oxjIxh3XxPyc_pvo' -f Dockerfile .

run:
	docker run $(IMAGE)

PHONY: build run