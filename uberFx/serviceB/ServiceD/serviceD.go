package ServiceD

import "fmt"

type ForthService struct {
}

func (f *ForthService) ForthServiceStart() {
	fmt.Println("Started forth service...")
}
