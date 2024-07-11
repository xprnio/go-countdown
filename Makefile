all: linux win

linux: linux_x64
linux_x64: bin/linux_x64/go-countdown
bin/linux_x64/go-countdown: cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux_x64/go-countdown cmd/main.go

win: win_x64
win_x64: bin/win_x64/go-countdown.exe
bin/win_x64/go-countdown.exe: cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/win_x64/go-countdown.exe cmd/main.go
