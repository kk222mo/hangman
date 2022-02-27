go mod tidy

cd data
go build
cd ..
cd database
go build
cd ..
cd server
go build
cd ..
go install github.com/kk222mo/hangman
