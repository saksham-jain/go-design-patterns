package notify_engine

type Subject interface {
	Register(o Observer)
	Deregister(o Observer)
	UpdateInventory(c int64)
}
