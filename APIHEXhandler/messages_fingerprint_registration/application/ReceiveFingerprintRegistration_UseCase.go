package application

type ReceiveFingerprintRegistrationUseCase struct {
    sf_uc *SendFingerprintRegistrationUseCase
}

func NewReceiveFingerprintRegistrationUseCase(sf_uc *SendFingerprintRegistrationUseCase) ReceiveFingerprintRegistrationUseCase {
    return ReceiveFingerprintRegistrationUseCase{sf_uc: sf_uc}
}

func (r *ReceiveFingerprintRegistrationUseCase) Execute(msg string) error {
    return r.sf_uc.Execute([]byte(msg))
}