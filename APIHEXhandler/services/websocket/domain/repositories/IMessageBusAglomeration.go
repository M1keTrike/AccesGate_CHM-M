package repositories

type IMessageBusAglomeration interface {
	Send(msg []byte) error
}