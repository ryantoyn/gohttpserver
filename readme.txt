go run C:\Users\suwatch\go\src\gohttpserver\main.go

curl 127.0.0.1:8080/foo
Hello, /foo!

cd /d C:\Users\suwatch\go\src\gohttpserver
go build
gohttpserver.exe

curl 127.0.0.1:8080/foo
Hello, /foo!

REM assuming GOPATH=C:\Users\suwatch\go
REM go install with project relative to %GOPATH%\src
REM output exe to %GOPATH%\bin
go install gohttpserver
C:\Users\suwatch\go\bin\gohttpserver.exe

docker build -t suwatch01/gohttpserver C:\Users\suwatch\go\src\gohttpserver
docker run --name gohttpserver --rm -it -p 8080:8080 suwatch01/gohttpserver
docker push suwatch01/gohttpserver:latest
