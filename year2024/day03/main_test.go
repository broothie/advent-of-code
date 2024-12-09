package main

import (
	"testing"

	"github.com/shoenig/test"
)

func Test_mulRegexp(t *testing.T) {
	test.True(t, mulRegexp.MatchString("mul(1,1)"))
	test.True(t, mulRegexp.MatchString("mul(22,22)"))
	test.True(t, mulRegexp.MatchString("mul(333,333)"))
	test.True(t, mulRegexp.MatchString("mul(1,333)"))

	test.False(t, mulRegexp.MatchString("mul(1 ,333)"))
	test.False(t, mulRegexp.MatchString("mul(1 , 333)"))
	test.False(t, mulRegexp.MatchString("mul( 1,333 )"))
}
