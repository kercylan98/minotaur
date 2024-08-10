go build -o protoc-gen-go-minotaur-vivid-typed.exe

protoc -I=./options --plugin=protoc-gen-minotaur-vivid-typed=protoc-gen-go-minotaur-vivid-typed.exe --go_out=./test --go_opt=paths=source_relative --go-minotaur-vivid-typed_out=./test --go-minotaur-vivid-typed_opt=paths=source_relative --proto_path=./test ./test/*.proto