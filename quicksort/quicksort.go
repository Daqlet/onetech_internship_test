package quicksort

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	pivot := a[len(a)-1]
	m := 0
	for i := 0; i < len(a); i++ {
		if pivot > a[i] {
			a[m], a[i] = a[i], a[m]
			m++
		}
	}
	a[m], a[len(a)-1] = a[len(a)-1], a[m]
	QuickSort(a[:m])
	QuickSort(a[m+1:])
}
