all: build clean
build: compile 
	docker build -t nohaapav/app .
compile:  
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . 
clean: 
	rm -rf main

