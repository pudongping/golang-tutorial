package redis_big_key

import (
	"fmt"
	"testing"
)

func TestWriteBigKey(t *testing.T) {
	WriteBigKey()
}

func TestScanBigKey(t *testing.T) {
	maxMemory := int64(10)
	keys := ScanBigKey(maxMemory)
	fmt.Println(keys)
	WriteKeysToFile(keys)
	ClearKeys(keys)
}
