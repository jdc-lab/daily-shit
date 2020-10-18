module github.com/jdc-lab/daily-shit/gateway

go 1.15

replace (
	github.com/jdc-lab/daily-shit/common => ../common
	github.com/jdc-lab/daily-shit/user-service => ../user-service
)

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/jdc-lab/daily-shit/common v0.0.0-20201017204306-a44e64b0653a
	github.com/jdc-lab/daily-shit/user-service v0.0.0-20201017185229-e8a64ff9e677
	github.com/opentracing/opentracing-go v1.1.0
	github.com/vektah/gqlparser/v2 v2.1.0
	google.golang.org/grpc v1.33.0
)
