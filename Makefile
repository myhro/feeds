CF_PAGES_WEBHOOK ?= http://httpbin.org/post
DIST_FOLDER = dist

.PHONY: autossegredos copasa liquipedia

clean:
	rm -rf $(DIST_FOLDER)/

autossegredos: dist
	go run main.go autossegredos > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia oldnewthing
	cp 404.html $(DIST_FOLDER)/

copasa: dist
	go run main.go copasa > $(DIST_FOLDER)/copasa.xml

deploy:
	@echo $(CF_PAGES_WEBHOOK) | xargs curl -X POST -s

dist:
	mkdir -p $(DIST_FOLDER)

golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

lint:
	golangci-lint run

liquipedia: dist
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

oldnewthing: dist
	go run main.go oldnewthing > $(DIST_FOLDER)/oldnewthing.xml

serve:
	npx wrangler pages dev $(DIST_FOLDER)/

test:
	go test -v ./...
