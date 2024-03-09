build/docker:
	docker build . -f ./deploy/Dockerfile -t egrant:api

run/docker:
	docker run --rm -p 8000:8000 egrant:api

run/api-dev:
	npm run build --prefix ./js
	cp -a ./js/dist/. ./public
	poetry run uvicorn api.main:app --reload

run/web-dev:
	npm run dev --prefix ./js

run/pocket:
	./pocket/pocketbase --dev serve

get/pocket:
	curl -JLO https://github.com/pocketbase/pocketbase/releases/download/v0.20.7/pocketbase_0.20.7_linux_amd64.zip
