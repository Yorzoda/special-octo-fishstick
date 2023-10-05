package serviceA

import (
	"fmt"
	"github.com/special-octo-fishstick/uberFx/serviceA/serviceC"
)

type FirsService struct {
	serviceC.ThirdService
}

func NewFirstSrv() *FirsService {
	return &FirsService{}
}

func (s *FirsService) FirstStartService() {

	fmt.Println("Started first service...")

}
