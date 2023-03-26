CF_PAGES_WEBHOOK ?= http://httpbin.org/post
CODE_FOLDERS = bin/ src/
DIST_FOLDER = dist

autossegredos: dist
	npx ts-node ./bin/autossegredos.ts > $(DIST_FOLDER)/autossegredos.xml

build: autossegredos liquipedia oldnewthing teamspeak
	cp 404.html $(DIST_FOLDER)/

check:
	npx prettier --check $(CODE_FOLDERS)

clean:
	rm -rf $(DIST_FOLDER)/

deploy:
	@echo $(CF_PAGES_WEBHOOK) | xargs curl -X POST -s

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

tsc:
	npx tsc --noEmit
