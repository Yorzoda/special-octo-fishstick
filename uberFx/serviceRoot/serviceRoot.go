package ServiceRoot

import (
	"fmt"
	"github.com/special-octo-fishstick/uberFx/serviceA"
	"github.com/special-octo-fishstick/uberFx/serviceB"
)

type RootService struct {
	first  *serviceA.FirsService
	second *serviceB.SecondService
}

func NewRootService(first *serviceA.FirsService, second *serviceB.SecondService) *RootService {
	return &RootService{first, second}
}

func (r *RootService) StartRootService() {
	fmt.Println("Started root service")
	ch := make(chan bool)
	r.first.FirstStartService(ch)
	r.second.SecondServerStart(ch)

}
