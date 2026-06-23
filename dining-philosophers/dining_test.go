package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i <= 4; i++ {
		orderFinished = []string{}

		dine()

		if len(orderFinished) != 5 {
			t.Error("Something is wrong", len(orderFinished))
		}
	}
}
