build:
	go build

test:
	@go test `go list ./... | grep -vF /vendor/`

lint:
	@gometalinter.v2 \
		--vendor \
		--deadline=120s \
		--disable=gotype \
		--disable=gas \
		--exclude=".*should have comment or be unexported.*" \
		./...

docserve:
	@echo "Starting GoDoc server on http://0.0.0.0:6060"
	godoc -http=:6060

run: build
	./dad

install:
	go install

build-release:
	go build -ldflags "-s -w"
	upx dad  # brew install upx

devup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install --update
	dep ensure
