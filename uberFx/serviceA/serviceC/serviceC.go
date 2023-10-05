package serviceC

import "fmt"

type ThirdService struct {
}

func NewThirdService() *ThirdService {
	return &ThirdService{}
}

func (t *ThirdService) ThirdServiceStart() {
	fmt.Println("Started third service...")
}
