package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_joltageForBank(t *testing.T) {
	type TestCase struct {
		bank            string
		expectedJoltage int
	}

	testCases := []TestCase{
		{bank: "987654321111111", expectedJoltage: 987654321111},
		{bank: "811111111111119", expectedJoltage: 811111111119},
		{bank: "234234234234278", expectedJoltage: 434234234278},
		{bank: "818181911112111", expectedJoltage: 888911112111},
	}

	for _, testCase := range testCases {
		t.Run(testCase.bank, func(t *testing.T) {
			bank, err := parseBank(testCase.bank)
			require.NoError(t, err)

			joltage, err := joltageForBank(bank)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedJoltage, joltage)
		})
	}
}
