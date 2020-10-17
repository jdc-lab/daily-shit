module github.com/jdc-lab/daily-shit/test-client

go 1.15

replace github.com/jdc-lab/daily-shit/user-service => ../user-service

require (
	github.com/jdc-lab/daily-shit/user-service v0.0.0-20201017185229-e8a64ff9e677
	google.golang.org/grpc v1.33.0
)
