module github.com/catcherwong/rest-api-sample

go 1.13

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.25.1

replace google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20191108220845-16a3f7862a1a

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis/v7 v7.0.0-beta.4
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/prometheus/client_golang v1.2.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	google.golang.org/grpc v1.23.0
)
