package task

import "nokowebapi/cores"

type DependsOnImpl interface {
	GetTarget() string
	GetWaiter() string
}

type DependsOn struct {
	Target string `mapstructure:"target" json:"target" yaml:"target"`
	Waiter string `mapstructure:"waiter" json:"waiter" yaml:"waiter"`
}

func NewDependsOn(target string, waiter string) DependsOnImpl {
	return &DependsOn{
		Target: target,
		Waiter: waiter,
	}
}

func (d *DependsOn) GetName() string {
	return "depends_on"
}

func (d *DependsOn) GetTarget() string {
	return d.Target
}

func (d *DependsOn) GetWaiter() string {
	return d.Waiter
}

type Impl interface{}

type Task struct {
	Name      string            `mapstructure:"name" json:"name" yaml:"name"`
	Execute   string            `mapstructure:"execute" json:"execute" yaml:"execute"`
	Args      []string          `mapstructure:"args" json:"args" yaml:"args"`
	Workdir   string            `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
	Environ   []string          `mapstructure:"environ" json:"environ" yaml:"environ"`
	Stdin     string            `mapstructure:"stdin" json:"stdin" yaml:"stdin"`
	Stdout    string            `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	Stderr    string            `mapstructure:"stderr" json:"stderr" yaml:"stderr"`
	Network   cores.NetworkImpl `mapstructure:"network" json:"network" yaml:"network"`
	DependsOn []DependsOnImpl   `mapstructure:"depends_on" json:"dependsOn" yaml:"depends_on"`
}

func (Task) GetName() string {
	return "task"
}

func New() Impl {
	return &Task{}
}
