package interface_one

import "fmt"

var (
	instance Countup
)

func GetCountUp() Countup {
	if instance == nil {
		instance = &RealCountup{}
	}
	return instance
}

type Countup interface {
	Click()
	UnClick()
}

type RealCountup struct {}

func (rc *RealCountup) Click() {
	fmt.Print("Click!\n")
}

func (rc *RealCountup) UnClick() {
	fmt.Print("Unclick!\n")
}

type MockCountup struct {
	Clicks int
}

func (mc *MockCountup) Click() {
	fmt.Print("Click!\n")
	mc.Clicks++
}

func (mc *MockCountup) UnClick() {
	fmt.Print("Unclick!\n")
	mc.Clicks--
}

