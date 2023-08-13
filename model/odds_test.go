package model

import (
	"fmt"
	"math"
	"testing"
)

func TestNormalDistribution(t *testing.T) {
	p := normalDistribution(0, 0, 1)
	if math.Abs(p-0.39894228) > 1e-8 {
		t.Errorf("TestNormalDistribution return error %v", p)
		return
	}
}

func TestCalculateOdds(t *testing.T) {
	for i := 0; i <= 50; i++ {
		odds, _ := CalculateOdds(float64(i)*1e9, 50*1e9, 60*1e9)
		fmt.Printf("%d -> %.2f\n", i, odds)
	}
}
