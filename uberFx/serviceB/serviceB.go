package serviceB

import (
	"fmt"
)

type SecondService struct {
}

func NewSecondSrv() *SecondService {
	return &SecondService{}
}

func (s *SecondService) SecondServerStart(ch chan bool) {
	fmt.Println("Started second service...")
	ch <- true
}
