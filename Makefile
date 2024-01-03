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