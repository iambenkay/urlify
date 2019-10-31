fuser -k 8000/tcp
go build -o ./bin/urlspace || exit 1
./bin/urlspace
