dactyl: dactyl.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dactyl .

.PHONY: clean docker

clean:
	rm -rf *~ dactyl

docker: dactyl
	docker build -t cleardata/dactyl:latest .
