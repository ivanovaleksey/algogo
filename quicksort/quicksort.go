package quicksort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func Sort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, l, r int) {
	//log.Printf("sort l=%d, r=%d\n", l, r)

	for l < r {
		{
			//swap random element with last for randomization
			idx := rand.Intn(r + 1 - l)
			//log.Printf("randomize %d\n", idx+l)
			swap(arr, idx+l, r)
			//log.Printf("%v // %v\n", arr, arr[l:r+1])
		}

		m := partition(arr, l, r)
		//log.Printf("m=%d\n", m)

		sort(arr, l, m-1)
		l = m + 1
	}
}

func partition(arr []int, l, r int) int {
	i := l - 1

	for j := l; j <= r; j++ {
		if arr[j] <= arr[r] {
			swap(arr, i+1, j)
			i++
		}
	}

	return i
}

func swap(arr []int, i, j int) {
	//log.Printf("swap i,j=%d,%d\n", i, j)
	if i == j {
		return
	}
	arr[i], arr[j] = arr[j], arr[i]
}
