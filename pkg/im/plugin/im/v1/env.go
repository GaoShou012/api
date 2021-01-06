package im_v1

import (
	"hash/crc32"
	"math"
	"runtime"
)

func ParallelUUID(uuid string) int {
	v := int(crc32.ChecksumIEEE([]byte(uuid)))
	return int(math.Abs(float64(v))) % runtime.NumCPU()
}
