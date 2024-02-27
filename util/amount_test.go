// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util_test

import (
	"github.com/seracoin/serad/domain/consensus/utils/constants"
	"math"
	"testing"

	. "github.com/seracoin/serad/util"
)

func TestAmountCreation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		valid    bool
		expected Amount
	}{
		// Positive tests.
		{
			name:     "zero",
			amount:   0,
			valid:    true,
			expected: 0,
		},
		{
			name:     "max producible",
			amount:   29e9,
			valid:    true,
			expected: Amount(constants.MaxSeep),
		},
		{
			name:     "one hundred",
			amount:   100,
			valid:    true,
			expected: 100 * constants.SeepPerSera,
		},
		{
			name:     "fraction",
			amount:   0.01234567,
			valid:    true,
			expected: 1234567,
		},
		{
			name:     "rounding up",
			amount:   54.999999999999943157,
			valid:    true,
			expected: 55 * constants.SeepPerSera,
		},
		{
			name:     "rounding down",
			amount:   55.000000000000056843,
			valid:    true,
			expected: 55 * constants.SeepPerSera,
		},

		// Negative tests.
		{
			name:   "not-a-number",
			amount: math.NaN(),
			valid:  false,
		},
		{
			name:   "-infinity",
			amount: math.Inf(-1),
			valid:  false,
		},
		{
			name:   "+infinity",
			amount: math.Inf(1),
			valid:  false,
		},
	}

	for _, test := range tests {
		a, err := NewAmount(test.amount)
		switch {
		case test.valid && err != nil:
			t.Errorf("%v: Positive test Amount creation failed with: %v", test.name, err)
			continue
		case !test.valid && err == nil:
			t.Errorf("%v: Negative test Amount creation succeeded (value %v) when should fail", test.name, a)
			continue
		}

		if a != test.expected {
			t.Errorf("%v: Created amount %v does not match expected %v", test.name, a, test.expected)
			continue
		}
	}
}

func TestAmountUnitConversions(t *testing.T) {
	tests := []struct {
		name      string
		amount    Amount
		unit      AmountUnit
		converted float64
		s         string
	}{
		{
			name:      "MSRA",
			amount:    Amount(constants.MaxSeep),
			unit:      AmountMegaSRA,
			converted: 29000,
			s:         "29000 MSRA",
		},
		{
			name:      "kSRA",
			amount:    44433322211100,
			unit:      AmountKiloSRA,
			converted: 444.33322211100,
			s:         "444.333222111 kSRA",
		},
		{
			name:      "SRA",
			amount:    44433322211100,
			unit:      AmountSRA,
			converted: 444333.22211100,
			s:         "444333.222111 SRA",
		},
		{
			name:      "mSRA",
			amount:    44433322211100,
			unit:      AmountMilliSRA,
			converted: 444333222.11100,
			s:         "444333222.111 mSRA",
		},
		{

			name:      "μSRA",
			amount:    44433322211100,
			unit:      AmountMicroSRA,
			converted: 444333222111.00,
			s:         "444333222111 μSRA",
		},
		{

			name:      "seep",
			amount:    44433322211100,
			unit:      AmountSeep,
			converted: 44433322211100,
			s:         "44433322211100 Seep",
		},
		{

			name:      "non-standard unit",
			amount:    44433322211100,
			unit:      AmountUnit(-1),
			converted: 4443332.2211100,
			s:         "4443332.22111 1e-1 SRA",
		},
	}

	for _, test := range tests {
		f := test.amount.ToUnit(test.unit)
		if f != test.converted {
			t.Errorf("%v: converted value %v does not match expected %v", test.name, f, test.converted)
			continue
		}

		s := test.amount.Format(test.unit)
		if s != test.s {
			t.Errorf("%v: format '%v' does not match expected '%v'", test.name, s, test.s)
			continue
		}

		// Verify that Amount.ToSRA works as advertised.
		f1 := test.amount.ToUnit(AmountSRA)
		f2 := test.amount.ToSRA()
		if f1 != f2 {
			t.Errorf("%v: ToSRA does not match ToUnit(AmountSRA): %v != %v", test.name, f1, f2)
		}

		// Verify that Amount.String works as advertised.
		s1 := test.amount.Format(AmountSRA)
		s2 := test.amount.String()
		if s1 != s2 {
			t.Errorf("%v: String does not match Format(AmountSRA): %v != %v", test.name, s1, s2)
		}
	}
}

func TestAmountMulF64(t *testing.T) {
	tests := []struct {
		name string
		amt  Amount
		mul  float64
		res  Amount
	}{
		{
			name: "Multiply 0.1 SRA by 2",
			amt:  100e5, // 0.1 SRA
			mul:  2,
			res:  200e5, // 0.2 SRA
		},
		{
			name: "Multiply 0.2 SRA by 0.02",
			amt:  200e5, // 0.2 SRA
			mul:  1.02,
			res:  204e5, // 0.204 SRA
		},
		{
			name: "Round down",
			amt:  49, // 49 Seep
			mul:  0.01,
			res:  0,
		},
		{
			name: "Round up",
			amt:  50, // 50 Seep
			mul:  0.01,
			res:  1, // 1 Seep
		},
		{
			name: "Multiply by 0.",
			amt:  1e8, // 1 SRA
			mul:  0,
			res:  0, // 0 SRA
		},
		{
			name: "Multiply 1 by 0.5.",
			amt:  1, // 1 Seep
			mul:  0.5,
			res:  1, // 1 Seep
		},
		{
			name: "Multiply 100 by 66%.",
			amt:  100, // 100 Seep
			mul:  0.66,
			res:  66, // 66 Seep
		},
		{
			name: "Multiply 100 by 66.6%.",
			amt:  100, // 100 Seep
			mul:  0.666,
			res:  67, // 67 Seep
		},
		{
			name: "Multiply 100 by 2/3.",
			amt:  100, // 100 Seep
			mul:  2.0 / 3,
			res:  67, // 67 Seep
		},
	}

	for _, test := range tests {
		a := test.amt.MulF64(test.mul)
		if a != test.res {
			t.Errorf("%v: expected %v got %v", test.name, test.res, a)
		}
	}
}
