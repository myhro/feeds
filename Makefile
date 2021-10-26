DIST_FOLDER = dist

build:
	mkdir -p $(DIST_FOLDER)
	go run main.go liquipedia > $(DIST_FOLDER)/liquipedia.xml

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	npx wrangler publish

deploy-prod:
	npx wrangler publish --env production

test:
	go test -v ./...
