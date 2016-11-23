package models

import "encoding/binary"

//itob converts an int to an 8-byte big endian encoded byte slice.
//In general, use big endian encoding for integers because it provides lexicographical sorting which is important when we want to iterate over our data.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
