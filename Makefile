CODE_FOLDERS = bin/ src/ tests/
DIST_FOLDER = dist

autossegredos: dist
	wget -O $(DIST_FOLDER)/autossegredos.xml https://www.autossegredos.com.br/category/segredos/feed/

build: autossegredos
	cp 404.html $(DIST_FOLDER)/

check:
	npx prettier --check $(CODE_FOLDERS)

check-updates:
	npx npm-check-updates -u -t minor

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	./deploy.sh

dist:
	mkdir -p $(DIST_FOLDER)

lint:
	DEBUG=eslint:cli-engine npx eslint $(CODE_FOLDERS)

liquipedia: dist
	npx ts-node ./bin/liquipedia.ts > $(DIST_FOLDER)/liquipedia.xml

oldnewthing: dist
	npx ts-node ./bin/oldnewthing.ts > $(DIST_FOLDER)/oldnewthing.xml

prettier:
	npx prettier --write $(CODE_FOLDERS)

serve:
	BROWSER=none npx wrangler pages dev $(DIST_FOLDER)/

sourcegraph: dist
	npx ts-node ./bin/sourcegraph.ts > $(DIST_FOLDER)/sourcegraph.xml

teamspeak: dist
	npx ts-node ./bin/teamspeak.ts > $(DIST_FOLDER)/teamspeak.xml

test:
	npx mocha -r ts-node/register src/**/*.test.ts

test-integration:
	npx mocha -r ts-node/register tests/integration.test.ts

tsc:
	npx tsc --noEmit
