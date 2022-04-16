//go:build !linux
// +build !linux

package kernel

import (
	"flashcat.cloud/categraf/inputs"
	"flashcat.cloud/categraf/types"
)

const InputName = "kernel"

type KernelStats struct {
	PrintConfigs    bool  `toml:"print_configs"`
	IntervalSeconds int64 `toml:"interval_seconds"`
}

func init() {
	inputs.Add(InputName, func() inputs.Input {
		return &KernelStats{}
	})
}

func (s *KernelStats) GetInputName() string {
	return InputName
}

func (s *KernelStats) GetIntervalSeconds() int64 {
	return s.IntervalSeconds
}

func (s *KernelStats) Init() error {
	return nil
}

func (s *KernelStats) Gather() (samples []*types.Sample) {
	return
}
