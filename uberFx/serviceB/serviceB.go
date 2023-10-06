package serviceB

import (
	"fmt"
	"github.com/special-octo-fishstick/uberFx/serviceB/ServiceD"
)

type SecondService struct {
	*ServiceD.ForthService
}

func NewSecondSrv(frSrv *ServiceD.ForthService) *SecondService {
	return &SecondService{frSrv}
}

func (s *SecondService) SecondServerStart() {
	fmt.Println("Started second service...")
}
