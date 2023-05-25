BINARY_NAME=examples
.DEFAULT_GOAL := run

build:
	go build -o ./bin/${BINARY_NAME} main.go

clean:
	go clean
	rm -rf bin

dep: vet
	go mod download

vet:
	go vet

run:
	make clean
	make dep
	make build
	nohup ./bin/${BINARY_NAME} > ${BINARY_NAME}.nohup.out  &
	tail -20f ${BINARY_NAME}.nohup.out
fast:
	make clean
	make dep
	make build
	./bin/${BINARY_NAME}