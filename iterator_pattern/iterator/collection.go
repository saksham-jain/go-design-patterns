package iterator

type Collection interface {
	Iterator() Iterator // this can be avoided and iterator instance can be directly created in iterator class.
}
