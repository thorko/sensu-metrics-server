build:
	env GOOS=linux GARCH=amd64 go build -o bin/getmetrics cmd/getmetrics.go

docker:
	docker build -t getmetrics:${VERSION} .
	docker run -d --name=getmetrics getmetrics:${VERSION}
	docker cp getmetrics:/tmp/getmetrics ./bin/getmetrics
	docker stop getmetrics
	docker rm getmetrics
	docker rmi --force getmetrics:${VERSION}

release:
	tar czf /tmp/sensu-metrics-server_${VERSION}_linux_amd64.tar.gz bin/ 
	sum=$$(sha512sum /tmp/sensu-metrics-server_${VERSION}_linux_amd64.tar.gz | cut -d ' ' -f 1); \
	f=$$(basename sensu-metrics-server_${VERSION}_linux_amd64.tar.gz); \
	echo $$sum $${f} > /tmp/sensu-metrics-server_${VERSION}_sha512_checksums.txt; \
	echo $$sum;
