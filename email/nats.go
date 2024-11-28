package email

import "github.com/nats-io/nats.go"

var (
	Nc  *nats.Conn
	err error
)

func initNats() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
}
