start:
	docker-compose up -d --build
test:
	docker exec -it checkout go test -tags testing -covermode=atomic ./core/...