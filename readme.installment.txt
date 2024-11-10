go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc --go_out=. --go-grpc_out=. proto/*.proto

cd book-service
book-service  go mod init github.com/FARRAS-DARKUNO/library-management/book-service
cd ../author-service
go mod init github.com/FARRAS-DARKUNO/library-management/author-service
cd ../category-service
go mod init github.com/FARRAS-DARKUNO/library-management/category-service
cd ../user-service
go mod init github.com/FARRAS-DARKUNO/library-management/user-service

cd book-service
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/go-redis/redis/v8
go get github.com/golang-jwt/jwt/v4
go get google.golang.org/grpc
go get google.golang.org/protobuf

go get github.com/jackc/pgx/v5/pgconn@v5.5.5
go get github.com/jackc/puddle/v2@v2.2.1
go get github.com/gofiber/fiber/v2@v2.52.5