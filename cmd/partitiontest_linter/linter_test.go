// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

<<<<<<<< HEAD:cmd/partitiontest_linter/linter_test.go
package linter
========
package logic
>>>>>>>> teal4-bench:data/transactions/logic/fields_test.go

import (
	"testing"

<<<<<<<< HEAD:cmd/partitiontest_linter/linter_test.go
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
========
	"github.com/stretchr/testify/require"
)

func TestArrayFields(t *testing.T) {
	require.Equal(t, len(TxnaFieldNames), len(TxnaFieldTypes))
	require.Equal(t, len(txnaFieldSpecByField), len(TxnaFieldTypes))
>>>>>>>> teal4-bench:data/transactions/logic/fields_test.go
}
