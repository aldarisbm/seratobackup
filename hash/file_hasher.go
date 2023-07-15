package hash

import (
	"hash"
	"os"
)

type Hasher interface {
	HashFile(file os.File) hash.Hash
}
