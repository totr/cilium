// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package metrics

import (
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MetricsSuite struct{}

var _ = Suite(&MetricsSuite{})

func (s *MetricsSuite) TestLastInteraction(c *C) {
	LastInteraction.Reset()
	c.Assert(time.Since(LastInteraction.Time()) < time.Second, Equals, true)
}
