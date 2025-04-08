package repositories

type IMessageBusAlarm interface {
    Send(msg []byte) error
}