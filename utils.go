package GrapeDB

import (
	"hash/fnv"
)

func Hash(name string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(name))
	return h.Sum32()
}
