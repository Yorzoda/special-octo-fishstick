package serviceA

import (
	"fmt"
	"github.com/special-octo-fishstick/uberFx/serviceA/serviceC"
	"github.com/special-octo-fishstick/uberFx/serviceB/ServiceD"
)

type FirsService struct {
	third  *serviceC.ThirdService
	fourth *ServiceD.ForthService
}

func NewFirstSrv(trdSrv *serviceC.ThirdService, forthSrv *ServiceD.ForthService) *FirsService {
	return &FirsService{trdSrv, forthSrv}
}

func (s *FirsService) FirstStartService(ch chan bool) {
	fmt.Println("Started first service...")
	go func() {
		_ = <-ch
		s.third.ThirdServiceStart()
		s.fourth.ForthServiceStart()
	}()
}
