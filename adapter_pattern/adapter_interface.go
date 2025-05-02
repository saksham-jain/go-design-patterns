package adapterpattern

type ChargerType string

const (
	TwoPin   ChargerType = "TwoPin"
	ThreePin ChargerType = "ThreePin"
)

type ChargerAdapter interface {
	Charge() error
}
