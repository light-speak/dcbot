module github.com/lty5240/dcbot/client-main

go 1.13

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20200130185559-910be7a94367

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.3.0

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/micro/go-grpc v1.0.1
)
