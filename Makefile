dev-run:
	docker-compose -f ./docker/dev.yml up -d
	goose up
dev-stop:
	 docker-compose -f ./docker/dev.yml down
migrations-up:
	goose up
migrations-down:
	goose down
set-env:
	source .env