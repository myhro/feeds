DIST_FOLDER = dist

autossegredos: dist
	wget -O $(DIST_FOLDER)/autossegredos.xml https://www.autossegredos.com.br/category/segredos/feed/

build: autossegredos
	cp 404.html $(DIST_FOLDER)/

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	./deploy.sh

dist:
	mkdir -p $(DIST_FOLDER)

serve:
	BROWSER=none npx wrangler pages dev $(DIST_FOLDER)/
