OUTPUT	=	simulation
FILES	=	*.go

all: build

build:
	go build -o $(OUTPUT) $(FILES)

run:
	./simulation map