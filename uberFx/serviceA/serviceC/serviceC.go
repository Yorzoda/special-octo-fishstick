package serviceC

import "fmt"

type ThirdService struct {
}

func (t *ThirdService) ThirdServiceStart() {
	fmt.Println("Started third service...")
}
