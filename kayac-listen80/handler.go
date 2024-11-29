package kayaclisten80

import "net/http"

type listen80Hander struct {
}

func NewHandler() *listen80Hander {
	return &listen80Hander{}
}

func (h *listen80Hander) Handle(req *http.Request) error {
	return nil
}
