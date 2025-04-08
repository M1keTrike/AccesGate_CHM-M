package repositories

type IMessageBusFingerprintRegistration interface {
    Send(msg []byte) error
}