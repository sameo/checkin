all: main.go
	go build -o checkin-server main.go

docker: all Dockerfile
	docker build -t checkin-server .

docker-run: docker
	docker run -p 9090:9090 checkin-server

clean:
	rm -rf checkin-server *~

