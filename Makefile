GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PLATFORM_FILES="./"
DIST_FILE="./golion"
DIST_FILE_LINUX="./golion_linux"

all: test build
build: 
	GO111MODULE=on $(GOBUILD) -o $(DIST_FILE) -v $(PLATFORM_FILES)
build-linux:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(DIST_FILE_LINUX) -v $(PLATFORM_FILES)
run:
	GO111MODULE=on $(GORUN) -v $(PLATFORM_FILES)
	./$(PLATFORM_FILES)
test: 
	$(GOTEST) -v ./...
clean: 
	GO111MODULE=on $(GOCLEAN) $(PLATFORM_FILES)
	rm -f $(DIST_FILE)
	rm -f $(DIST_FILE_LINUX)	
    
    
