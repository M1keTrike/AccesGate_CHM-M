package application

type ReceiveAlarmUseCase struct {
    sa_uc *SendAlarmUseCase
}

func NewReceiveAlarmUseCase(sa_uc *SendAlarmUseCase) ReceiveAlarmUseCase {
    return ReceiveAlarmUseCase{sa_uc: sa_uc}
}

func (r *ReceiveAlarmUseCase) Execute(msg string) error {
    return r.sa_uc.Execute([]byte(msg))
}