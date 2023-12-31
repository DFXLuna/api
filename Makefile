build/docker:
	docker build . -f ./deploy/Dockerfile -t egrant:api

dev:
	npm run build --prefix ./js
	cp -a ./js/dist/. ./public
	poetry run uvicorn api.main:app --reload