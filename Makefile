build:
	env CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bin/getmetrics cmd/getmetrics.go

release:
	tar czf /tmp/sensu-metrics-server_${VERSION}_linux_amd64.tar.gz bin/ 
	sum=$$(sha512sum /tmp/sensu-metrics-server_${VERSION}_linux_amd64.tar.gz | cut -d ' ' -f 1); \
	f=$$(basename sensu-metrics-server_${VERSION}_linux_amd64.tar.gz); \
	echo $$sum $${f} > /tmp/sensu-metrics-server_${VERSION}_sha512_checksums.txt; \
	echo $$sum;
