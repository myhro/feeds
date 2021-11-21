CF_PAGES_WEBHOOK ?= http://httpbin.org/post
DIST_FOLDER = dist

.PHONY: autossegredos copasa liquipedia

clean:
	rm -rf $(DIST_FOLDER)/

autossegredos: dist
	go run main.go autossegredos > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia
	cp 404.html $(DIST_FOLDER)/

copasa: dist
	go run main.go copasa > $(DIST_FOLDER)/copasa.xml

deploy:
	@echo $(CF_PAGES_WEBHOOK) | xargs curl -X POST

dist:
	mkdir -p $(DIST_FOLDER)

liquipedia: dist
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

serve:
	npx wrangler pages dev $(DIST_FOLDER)/

test:
	go test -v ./...
