package repositories

type IMessageBusFingerprint interface {
	Send(msg []byte) error
}