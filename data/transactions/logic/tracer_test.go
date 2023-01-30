// Copyright (C) 2019-2023 Algorand, Inc.
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

package logic_test

import (
	"testing"

	"github.com/algorand/go-algorand/data/basics"
	. "github.com/algorand/go-algorand/data/transactions/logic"
	"github.com/algorand/go-algorand/data/transactions/logic/mocktracer"
	"github.com/algorand/go-algorand/test/partitiontest"
	"github.com/stretchr/testify/require"
)

type tracerTestCase struct {
	name           string
	program        string
	evalProblems   []string
	expectedEvents []mocktracer.Event
}

func getSimpleTracerTestCases(mode RunMode) []tracerTestCase {
	return []tracerTestCase{
		{
			name:    "approve",
			program: debuggerTestProgramApprove,
			expectedEvents: mocktracer.FlattenEvents([][]mocktracer.Event{
				{
					mocktracer.BeforeProgram(mode),
				},
				mocktracer.OpcodeEvents(35, false),
				{
					mocktracer.AfterProgram(mode, false),
				},
			}),
		},
		{
			name:         "reject",
			program:      debuggerTestProgramReject,
			evalProblems: []string{"REJECT"},
			expectedEvents: mocktracer.FlattenEvents([][]mocktracer.Event{
				{
					mocktracer.BeforeProgram(mode),
				},
				mocktracer.OpcodeEvents(36, false),
				{
					mocktracer.AfterProgram(mode, false),
				},
			}),
		},
		{
			name:         "error",
			program:      debuggerTestProgramError,
			evalProblems: []string{"err opcode executed"},
			expectedEvents: mocktracer.FlattenEvents([][]mocktracer.Event{
				{
					mocktracer.BeforeProgram(mode),
				},
				mocktracer.OpcodeEvents(36, true),
				{
					mocktracer.AfterProgram(mode, true),
				},
			}),
		},
	}
}

func TestEvalWithTracer(t *testing.T) {
	partitiontest.PartitionTest(t)
	t.Parallel()

	t.Run("logicsig", func(t *testing.T) {
		t.Parallel()
		testCases := getSimpleTracerTestCases(ModeSig)
		for _, testCase := range testCases {
			testCase := testCase
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()
				mock := mocktracer.Tracer{}
				ep := DefaultEvalParams()
				ep.Tracer = &mock
				TestLogic(t, testCase.program, AssemblerMaxVersion, ep, testCase.evalProblems...)

				require.Equal(t, testCase.expectedEvents, mock.Events)
			})
		}
	})

	t.Run("simple app", func(t *testing.T) {
		t.Parallel()
		testCases := getSimpleTracerTestCases(ModeApp)
		for _, testCase := range testCases {
			testCase := testCase
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()
				mock := mocktracer.Tracer{}
				ep := DefaultEvalParams()
				ep.Tracer = &mock
				TestApp(t, testCase.program, ep, testCase.evalProblems...)

				require.Equal(t, testCase.expectedEvents, mock.Events)
			})
		}
	})

	t.Run("app with inner txns", func(t *testing.T) {
		t.Parallel()
		scenarios := mocktracer.GetTestScenarios()
		for name, makeScenario := range scenarios {
			makeScenario := makeScenario
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				mock := mocktracer.Tracer{}
				ep, tx, ledger := MakeSampleEnv()
				ep.Tracer = &mock

				// Establish FirstTestID as the app id, and fund it. We do this so that the created
				// inner app will get a sequential ID, which is what the mocktracer scenarios expect
				createdAppIndex := basics.AppIndex(FirstTestID)
				ledger.NewApp(tx.Receiver, createdAppIndex, basics.AppParams{})
				ledger.NewAccount(createdAppIndex.Address(), 200_000)
				tx.ApplicationID = createdAppIndex

				scenario := makeScenario(mocktracer.TestScenarioInfo{
					CallingTxn:   *tx,
					CreatedAppID: createdAppIndex,
				})

				var evalProblems []string
				switch scenario.Outcome {
				case mocktracer.RejectionOutcome:
					evalProblems = []string{"REJECT"}
				case mocktracer.ErrorOutcome:
					if scenario.ExpectedError == "overspend" {
						// the logic test ledger uses this error instead
						evalProblems = []string{"insufficient balance"}
					} else {
						evalProblems = []string{scenario.ExpectedError}
					}
				}

				ops := TestProg(t, scenario.Program, AssemblerNoVersion)
				TestAppBytes(t, ops.Program, ep, evalProblems...)

				// trim BeforeTxn and AfterTxn events from scenario.ExpectedEvents, since they are
				// not emitted from TestAppBytes
				require.Equal(t, scenario.ExpectedEvents[0].Type, mocktracer.BeforeTxnEvent)
				require.Equal(t, scenario.ExpectedEvents[len(scenario.ExpectedEvents)-1].Type, mocktracer.AfterTxnEvent)
				trimmedExpectedEvents := scenario.ExpectedEvents[1 : len(scenario.ExpectedEvents)-1]
				require.Equal(t, trimmedExpectedEvents, mock.Events)
			})
		}
	})
}
