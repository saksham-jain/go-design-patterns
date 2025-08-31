package pcbuilder

type PcBuilder interface {
	SetCpu(cpu string) PcBuilder
	SetGpu(gpu string) PcBuilder
	SetRam(ram string) PcBuilder
	Build() PcBuilder
}
