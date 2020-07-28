run-all-services:
	docker-compose up allinone

run-api-service:
	docker-compose up -d api

run-csv-service:
	docker-compose up  csv

new-migration:
	docker run --rm -v $(PWD)/schema:/db amacneil/dbmate new ${NAME}
