package repositories

type IMessageBusNFC interface {
	Send(msg []byte) error
}
