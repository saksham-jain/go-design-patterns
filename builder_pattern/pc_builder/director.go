package pcbuilder

type Director struct {
	PcBuilder PcBuilder
}

func NewDirector(pb PcBuilder) Director {
	return Director{
		PcBuilder: pb,
	}
}

func (d *Director) BuildGamingPc() PcBuilder {
	return d.PcBuilder.
		SetCpu("Intel i9").
		SetGpu("NVIDIA RTX 3090").
		SetRam("64GB").
		Build()
}

func (d *Director) BuildOfficePc() PcBuilder {
	return d.PcBuilder.
		SetCpu("Intel i5").
		SetGpu("Integrated Graphics").
		SetRam("16GB").
		Build()
}
