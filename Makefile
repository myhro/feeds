DIST_FOLDER = dist

.PHONY: autossegredos copasa liquipedia

clean:
	rm -rf $(DIST_FOLDER)/

autossegredos: dist
	go run main.go autossegredos > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia

copasa: dist
	go run main.go copasa > $(DIST_FOLDER)/copasa.xml

deploy:
	npx wrangler publish

deploy-prod:
	npx wrangler publish --env production

dist:
	mkdir -p $(DIST_FOLDER)

liquipedia: dist
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

test:
	go test -v ./...
