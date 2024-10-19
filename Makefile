FOCUS ?= .
ginkgo:
	CGO_ENABLED=1 ginkgo -r -v --race --trace --coverpkg=$(COVERPKG) --focus $(FOCUS) --skip=$(SKIP) --coverprofile=cover.out

go-cover: ginkgo
	@go tool cover -html=cover.out -o cover.html
	@go tool cover -func=cover.out
