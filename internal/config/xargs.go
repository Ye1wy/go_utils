package config

import "flag"

type ParallelFlag struct {
	parallel *int
}

func (p *ParallelFlag) Register() {
	p.parallel = flag.Int("parallel", 2, "Number of commands to run in parallel")
}

func (p *ParallelFlag) ContainedValue() interface{} {
	return *p.parallel
}

func (p *ParallelFlag) Name() string {
	return "parallel"
}
