package interface_one

import (
	"testing"
	"fmt"
)

func TestClick(t *testing.T) {
	mockCountup := MockCountup{}
	instance = &mockCountup
	GetCountUp().Click()
	fmt.Printf("My instance: %v and clicks %d\nInstance: %v", mockCountup, mockCountup.Clicks, instance)
	GetCountUp().Click()
	fmt.Printf("My instance: %v and clicks %d\nInstance: %v", mockCountup, mockCountup.Clicks, instance)
	GetCountUp().UnClick()
	fmt.Printf("My instance: %v and clicks %d\nInstance: %v", mockCountup, mockCountup.Clicks, instance)
	GetCountUp().UnClick()
	fmt.Printf("My instance: %v and clicks %d\nInstance: %v", mockCountup, mockCountup.Clicks, instance)

}
