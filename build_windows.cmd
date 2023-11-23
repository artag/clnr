cd src
echo Running tests:
go test ./...
echo Building application
go build -ldflags="-s -w" -o ../bin/clnr.exe
cd ..
echo Done