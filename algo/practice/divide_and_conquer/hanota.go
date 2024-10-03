package divideandconquer

import "container/list"

func move(src, tar *list.List) {
	e := src.Back()
	src.Remove(e)
	tar.PushBack(e.Value)
}

func dfsHanota(i int, src, buf, tar *list.List) {
	if i == 1 {
		move(src, tar)
		return
	}
	dfsHanota(i-1, src, tar, buf)
	move(src, tar)
	dfsHanota(i-1, buf, src, tar)
}

func solveHanota(A, B, C *list.List) {
	n := A.Len()
	dfsHanota(n, A, B, C)
}
