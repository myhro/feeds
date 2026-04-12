DIST_FOLDER = dist

autossegredos: dist
	./bin/fetch.sh $(DIST_FOLDER)/autossegredos.xml https://www.autossegredos.com.br/category/segredos/feed/

build: autossegredos
	cp 404.html $(DIST_FOLDER)/

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	./bin/deploy.sh

dist:
	mkdir -p $(DIST_FOLDER)

serve:
	BROWSER=none npx wrangler dev

staging: build
	npx wrangler versions upload --preview-alias staging
