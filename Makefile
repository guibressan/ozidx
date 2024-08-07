VERSION := "v0.1.0"

all:
	@mkdir -p out
	@CGO_ENABLED=0 \
		go build \
		-o out/ozidx \
		-ldflags="-s -w -X main.version=$(VERSION)" \
		main.go

clean:
	@rm -rf out
