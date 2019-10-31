fuser -k 8000/tcp
go build -o ./build/urlify || exit 1
./build/urlify
