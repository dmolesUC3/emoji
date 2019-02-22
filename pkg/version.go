package emoji

import (
	. "github.com/dmolesUC3/emoji/pkg/data"
	. "github.com/dmolesUC3/emoji/pkg/properties"
	"unicode"
)

// Version represents an Emoji major release, e.g. V5 for Emoji version 5.0
type Version int

const (
	V1 Version = 1
	V2 Version = 2
	V3 Version = 3
	V4 Version = 4
	V5 Version = 5
	// Starting at V11, Emoji version = Unicode version
	V11 Version = 11
	V12 Version = 12

	Latest = V12
)

// AllVersions lists all versions in order
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}

// HasFile returns true if this version has a file of the specified type, false
// otherwise. E.g., ZWJ (zero width joiner) sequences were introduced only in
// Emoji version 2.0, test files in version 4.0, and variation sequences in version
// 5.0.
func (v Version) HasFile(t FileType) bool {
	return v.FileBytes(t) != nil
}

// FileBytes returns the byte data of the Unicode.org source file of the specified type
// for this version, e.g. V12.FileBytes(Sequences) returns the contents of the file
// http://unicode.org/Public/emoji/12.0/emoji-sequences.txt
func (v Version) FileBytes(t FileType) []byte {
	if fileBytesByVersion, ok := fileBytesByVersionAndType[v]; ok {
		if bytes, ok := fileBytesByVersion[t]; ok {
			return bytes
		}
	}
	return nil
}

// RangeTable returns the Unicode range table for characters with the specified property
// in this Emoji version. Note that the range table reflects the ranges as defined in the
// source files from Unicode.org; ranges are guaranteed not to overlap, as per the RangeTable
// docs, but adjacent ranges are not coalesced.
func (v Version) RangeTable(property Property) *unicode.RangeTable {
	var exists bool
	var rtsByProperty map[Property]*unicode.RangeTable
	if rtsByProperty, exists = rangeTables[v]; !exists {
		rtsByProperty = map[Property]*unicode.RangeTable{}
		rangeTables[v] = rtsByProperty
	}
	var rt *unicode.RangeTable
	if rt, exists = rtsByProperty[property]; !exists {
		rt = ParseRangeTable(property, v.FileBytes(Data))
		rtsByProperty[property] = rt
	}
	return rt
}

// ------------------------------------------------------------
// Unexported symbols

var rangeTables = map[Version]map[Property]*unicode.RangeTable{}
