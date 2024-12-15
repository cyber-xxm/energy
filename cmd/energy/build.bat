SET CGO_ENABLED=0

set GOARCH=386
set GOOS=windows
go build -trimpath -ldflags "-s -w" -o energy-windows32.exe energy.go
md5.exe energy-windows32.exe

set GOARCH=amd64
set GOOS=windows
go build -trimpath -ldflags "-s -w" -o energy-windowsx64.exe energy.go
md5.exe energy-windowsx64.exe

set GOARCH=amd64
set GOOS=darwin
go build -trimpath -ldflags "-s -w" -o energy-macosx64 energy.go
md5.exe energy-macosx64

set GOARCH=arm64
set GOOS=darwin
go build -trimpath -ldflags "-s -w" -o energy-macosarm64 energy.go
md5.exe energy-macosarm64

set GOARCH=386
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linux32 energy.go
md5.exe energy-linux32

set GOARCH=amd64
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linuxx64 energy.go
md5.exe energy-linuxx64

set GOARCH=arm64
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linuxarm64 energy.go
md5.exe energy-linuxarm64

set GOARCH=arm
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linuxarm energy.go
md5.exe energy-linuxarm

pause