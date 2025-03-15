package repositories

type IMessageBusVAccess interface {
	Send(msg []byte) error
}
