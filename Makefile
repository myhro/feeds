DIST_FOLDER = dist

autossegredos: dist
	wget -O $(DIST_FOLDER)/autossegredos.xml https://www.autossegredos.com.br/category/segredos/feed/

build: clean autossegredos
	cp 404.html $(DIST_FOLDER)/

clean:
	rm -rf $(DIST_FOLDER)/

deploy: build
	./deploy.sh

dist:
	mkdir -p $(DIST_FOLDER)

serve:
	BROWSER=none npx wrangler dev

staging: build
	npx wrangler versions upload --preview-alias staging
