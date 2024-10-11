package chainofresponsibility

type DepartMent interface {
	execute(*Patient)
	setNext(DepartMent)
}
