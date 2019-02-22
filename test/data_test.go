package test

import (
	. "github.com/dmolesUC3/emoji/pkg"
	. "github.com/dmolesUC3/emoji/pkg/properties"
	. "gopkg.in/check.v1"
	"unicode"
)

type DataSuite struct {
}

var _ = Suite(&DataSuite{})

// Sample of emoji newly introduced by version
var samplesByPropertyAndVersion = map[Property]map[Version]string{
	Emoji: {
		V1:  "😀😃😄",	// 1F600, 1F603, 1F604
		V2:  "🗨",		// 1F5E8
		V3:  "🤣🤥🤤",	// 1F923, 1F925, 1F924
		V4:  "♀♂⚕",	// 2640, 2642, 2695
		V5:  "🤩🤪🤭",	// 1F929, 1F92A, 1F92D
		V11: "🥰🥵🥶",	// 1F970, 1F975, 1F976
		V12: "🥱🤎🤍",	// 1F971, 1F90E, 1F90D
	},
}

// Combined sample of specified version and all versions below it
func combinedSample(prop Property, v Version) string {
	samples := samplesByPropertyAndVersion[prop]
	sample := ""
	for _, v2 := range AllVersions {
		if v2 >= v {
			break
		}
		sample += samples[v2]
	}
	return sample
}

func (s *DataSuite) TestRangeTables(c *C) {
	ok := true
	for prop := range samplesByPropertyAndVersion {
		for _, v := range AllVersions {
			rt := v.RangeTable(prop)
			sample := combinedSample(prop, v)
			for _, r := range sample {
				inRange := unicode.In(r, rt)
				ok = ok && c.Check(inRange, Equals, true, Commentf("expected %v (%X) to be in %v range for %v, but was not", string(r), prop, v, r))
			}
		}
	}
	c.Assert(ok, Equals, true)
}
