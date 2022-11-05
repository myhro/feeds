CF_PAGES_WEBHOOK ?= http://httpbin.org/post
DIST_FOLDER = dist

.PHONY: liquipedia

autossegredos: dist
	npx ts-node ./src/autossegredos.ts > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia oldnewthing sourcegraph teamspeak
	cp 404.html $(DIST_FOLDER)/

check:
	npx prettier --check src/

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	@echo $(CF_PAGES_WEBHOOK) | xargs curl -X POST -s

dist:
	mkdir -p $(DIST_FOLDER)

eslint:
	DEBUG=eslint:cli-engine npx eslint src/

golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0

lint:
	golangci-lint run

liquipedia: dist
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

mocha:
	npx mocha -r ts-node/register src/**/*.test.ts

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
