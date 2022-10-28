package util

var Byte = newUtilByte()

type utilByte struct {
}

func newUtilByte() *utilByte {
	return &utilByte{}
}

// IsLetterUpper checks whether the given byte b is in upper case.
func (ub *utilByte) IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsLetterLower checks whether the given byte b is in lower case.
func (ub *utilByte) IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsLetter checks whether the given byte b is a letter.
func (ub *utilByte) IsLetter(b byte) bool {
	return ub.IsLetterUpper(b) || ub.IsLetterLower(b)
}
