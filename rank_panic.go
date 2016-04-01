package main
// under go1.6
import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Ad struct {
	Weight int
}

type AdsOrder []*Ad

func (a AdsOrder) Len() int {
	return len(a)
}

// 这么做排序可能不太好，但是暂时够用了
func (a AdsOrder) Less(i, j int) bool {
	if rand.Intn(2) == 0 {
		return true
	} else {
		return false
	}
	if a[i].Weight > a[j].Weight {
		return true
	}

	return false
}

func (a AdsOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	ads := []*Ad{
		&Ad{
			Weight: 10,
		},
		&Ad{
			Weight: 19,
		},
		&Ad{
			Weight: 1,
		},
		&Ad{
			Weight: 101,
		},
		&Ad{
			Weight: 4,
		},
		&Ad{
			Weight: 190,
		},
		&Ad{
			Weight: 7,
		},
		&Ad{
			Weight: 11,
		},
		&Ad{
			Weight: 98,
		},

		&Ad{
			Weight: 1,
		},
		&Ad{
			Weight: 9,
		},
		&Ad{
			Weight: 22,
		},
		&Ad{
			Weight: 11,
		},
		&Ad{
			Weight: 12,
		},
	}

	sort.Sort(AdsOrder(ads))
}
