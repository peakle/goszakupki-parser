COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)
GOOS=linux
GOARCH=amd64

NAME=loat

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -f $(NAME)

build: clean
	@echo ">> building..."
	@echo "Commit: $(COMMIT_ID)"
	@echo "Version: $(VERSION)"
	@ CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	    go build -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" -o $(NAME) ./cmd/main.go
	@chmod +x $(NAME)
