package ServiceD

import "fmt"

type ForthService struct {
}

func NewForthService() *ForthService {
	return &ForthService{}
}

func (f *ForthService) ForthServiceStart() {
	fmt.Println("Started forth service...")
}
