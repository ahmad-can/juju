// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package arch_test

import (
	jc "github.com/juju/testing/checkers"
	gc "launchpad.net/gocheck"

	"launchpad.net/juju-core/juju/arch"
	"launchpad.net/juju-core/testing/testbase"
)

type archSuite struct {
	testbase.LoggingSuite
}

var _ = gc.Suite(&archSuite{})

func (s *archSuite) TestHostArch(c *gc.C) {
	a, err := arch.HostArch()
	c.Assert(err, gc.IsNil)
	c.Assert(arch.IsSupportedArch(a), jc.IsTrue)
}

func (s *archSuite) TestNormaliseArch(c *gc.C) {
	for _, test := range []struct {
		raw  string
		arch string
	}{
		{"invalid", ""},
		{"amd64", "amd64"},
		{"x86_64", "amd64"},
		{"i386", "i386"},
		{"i486", "i386"},
		{"armv", "arm"},
		{"armv7", "arm"},
		{"aarch64", "arm64"},
		{"ppc64el", "ppc64"},
		{"ppc64le", "ppc64"},
	} {
		arch, err := arch.NormaliseArch(test.raw)
		if test.arch == "" {
			c.Check(err, gc.ErrorMatches, "unrecognised architecture:.*")
		} else {
			c.Check(arch, gc.Equals, test.arch)
		}
	}
}

func (s *archSuite) TestIsSupportedArch(c *gc.C) {
	for _, a := range arch.AllSupportedArches {
		c.Assert(arch.IsSupportedArch(a), jc.IsTrue)
	}
	c.Assert(arch.IsSupportedArch("invalid"), jc.IsFalse)
}
