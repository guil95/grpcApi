start:
	docker-compose up -d --build
test:
	go test -tags testing -covermode=atomic ./core/...