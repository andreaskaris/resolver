bin:
	mkdir bin

resolver: bin
	go build -o bin/resolver

#go-resolver: bin
#	GODEBUG=netdns=go+1 go build -o bin/go-resolver 

#cgo-resolver: bin
#	GODEBUG=netdns=cgo+1 go build -o bin/cgo-resolver

all: resolver #go-resolver cgo-resolver

.PHONY: clean
clean:
	rm -f bin/*
