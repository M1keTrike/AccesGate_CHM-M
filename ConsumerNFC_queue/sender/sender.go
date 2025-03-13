package sender

type Sender struct {
	url string
}

func NewSender(url string) *Sender {
	return &Sender{
		url: url,
	}
}

func (s *Sender) Send(body []byte) {

	//http.Post(s.url, "application/json", bytes.NewBuffer(body))
}
