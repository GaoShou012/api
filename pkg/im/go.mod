module im

go 1.15

require (
	framework v0.0.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/prometheus/client_golang v1.9.0 // indirect
)

replace framework => ./../framework
