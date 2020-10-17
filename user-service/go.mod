module github.com/jdc-lab/daily-shit/user-service

go 1.15

replace github.com/jdc-lab/daily-shit/common => ../common

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/hashicorp/consul v1.8.4 // indirect
	github.com/jdc-lab/daily-shit/common v0.0.0-20201017201424-0074e94ac05d
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	google.golang.org/genproto v0.0.0-20201015140912-32ed001d685c // indirect
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0
)
