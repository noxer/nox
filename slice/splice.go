package slice

// Splice removes a section of a slice and replaces it (optionally) with an
// insert. A length < 0 extends the section to remove to the end of the
// slice. Is the length longer than the slice it's trimmed. The index must be
// in bounds of the slice, otherwise Splice panics.
func Splice[T any](slice []T, index, length int, insert ...T) []T {
	// check if the index is out of bounds
	if index < 0 || index >= len(slice) {
		panic("slice index out of bounds")
	}

	// check if the length is negative
	if length < 0 {
		length = (len(slice) + length + 1) - index
	}

	// trim the length to fit the slice
	if index+length > len(slice) {
		// length reaches over the end of the slice
		// trim the length
		length = len(slice) - index
	}

	if length == len(insert) {
		// fastpath, we just copy the insert
		copy(slice[index:], insert)
		return slice
	}

	if length > len(insert) {
		// length is greater than len(insert)
		// we can copy the part to insert and move the remaining slice
		copy(slice[index:], insert)
		copy(slice[index+len(insert):], slice[index+length:])
		return slice[:len(slice)-(length-len(insert))]
	}

	// make space for the additional bytes
	slice = MakeRoom(slice, len(insert)-length)

	// move the last part of the slice out of the way
	copy(slice[index+len(insert)-length:], slice[index:])

	// copy the new part in
	copy(slice[index:], insert)

	return slice
}
