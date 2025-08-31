package pc

import pcbuilder "go-design-patterns/builder_pattern/pc_builder"

type Pc struct {
	Cpu string
	Gpu string
	Ram string
}

func NewPcBuilder() pcbuilder.PcBuilder {
	return &Pc{}
}

func (p *Pc) SetCpu(cpu string) pcbuilder.PcBuilder {
	p.Cpu = cpu
	return p
}

func (p *Pc) SetGpu(gpu string) pcbuilder.PcBuilder {
	p.Gpu = gpu
	return p
}

func (p *Pc) SetRam(ram string) pcbuilder.PcBuilder {
	p.Ram = ram
	return p
}

func (p *Pc) Build() pcbuilder.PcBuilder {
	return p
}
