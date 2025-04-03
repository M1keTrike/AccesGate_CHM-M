package application

type ReceiveFingerprintUseCase struct {
	sf_uc *SendFingerprintUseCase
}

func NewReceiveFingerprintUseCase(sf_uc *SendFingerprintUseCase) ReceiveFingerprintUseCase {
	return ReceiveFingerprintUseCase{sf_uc: sf_uc}
}

func (r *ReceiveFingerprintUseCase) Execute(msg string) error {
	return r.sf_uc.Execute([]byte(msg))
}