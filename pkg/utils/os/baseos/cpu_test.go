package baseos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetcpuinfo(t *testing.T) {
	var osobj BaseOS
	assert.NotNil(t, osobj.GetCPUInfo())
}
