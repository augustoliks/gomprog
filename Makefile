build-test:
	goreleaser --snapshot --skip-publish --rm-dist 

test:
	go test ./...
