swagger:
	rm -rf pkg/client pkg/models
	swagger generate client -f swagger/vro.json -t pkg

test:
	go build -o sdk-test