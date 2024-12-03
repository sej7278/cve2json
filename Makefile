$(shell mkdir -p builds/macos builds/linux builds/windows)

LDFLAGS=-ldflags "-s -w"

all: linux macos windows

linux:
	GOOS=linux GOARCH=amd64 go build -trimpath -buildmode=pie $(LDFLAGS) -o builds/linux/cve2json

macos:
	GOOS=darwin GOARCH=arm64 go build -trimpath -buildmode=pie $(LDFLAGS) -o builds/macos/cve2json

windows:
	GOOS=windows GOARCH=amd64 go build -trimpath -buildmode=pie $(LDFLAGS) -o builds/windows/cve2json.exe

clean:
	rm -rf builds/*

.PHONY: all clean linux macos windows
