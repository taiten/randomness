package detect

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestFactoryDetectFast(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	hit := "通过"
	pass, err := FactoryDetectFast(rand.Reader)
	if err != nil {
		hit = err.Error()
	}
	fmt.Printf("15种算法 出廠自检 1000组 10^6 bit: %v, hit: %s\n", pass, hit)
}
