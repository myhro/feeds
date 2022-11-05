CF_PAGES_WEBHOOK ?= http://httpbin.org/post
DIST_FOLDER = dist

.PHONY: autossegredos copasa liquipedia

autossegredos: dist
	go run main.go autossegredos > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia oldnewthing sourcegraph teamspeak
	cp 404.html $(DIST_FOLDER)/

check:
	npx prettier --check src/

clean:
	rm -rf $(DIST_FOLDER)/

copasa: dist
	go run main.go copasa > $(DIST_FOLDER)/copasa.xml

deploy:
	@echo $(CF_PAGES_WEBHOOK) | xargs curl -X POST -s

dist:
	mkdir -p $(DIST_FOLDER)

generate: oldnewthing sourcegraph teamspeak
	go run main.go generate -f autossegredos,liquipedia
	cp 404.html $(DIST_FOLDER)/

golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0

lint:
	golangci-lint run

liquipedia: dist
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

oldnewthing: dist
	npx ts-node ./src/oldnewthing.ts > $(DIST_FOLDER)/oldnewthing.xml

prettier:
	npx prettier --write src/

serve:
	npx wrangler pages dev $(DIST_FOLDER)/

sourcegraph: dist
	npx ts-node ./src/sourcegraph.ts > $(DIST_FOLDER)/sourcegraph.xml

teamspeak: dist
	npx ts-node ./src/teamspeak.ts > $(DIST_FOLDER)/teamspeak.xml

test:
	go test -v ./...

tsc:
	npx tsc --noEmit
