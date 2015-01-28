build: cmftp

cmftp: *.go
	go build -o $@ $^
