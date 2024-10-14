package visitor

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "rectangle"
}
