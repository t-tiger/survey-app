.PHONY: build-and-start
build-and-start:
	cd server && \
	make docker-build && \
	docker-compose down && \
	docker-compose up -d db && \
	sleep 6 && \
	docker-compose up -d survey-server && \
	cd ../client && \
	make docker-build && \
	docker-compose down && \
	docker-compose up -d survey-client

.PHONY: remove-containers
remove-containers:
	cd server && \
	docker-compose down && \
	cd ../client && \
	docker-compose down
