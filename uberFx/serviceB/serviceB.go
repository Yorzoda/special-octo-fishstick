package serviceB

import (
	"fmt"
	"github.com/special-octo-fishstick/uberFx/serviceB/ServiceD"
)

type SecondService struct {
	ServiceD.ForthService
}

func NewSecondSrv() *SecondService {
	return &SecondService{}
}

func (s *SecondService) SecondServerStart() {
	fmt.Println("Started second service...")
}
