package model

import (
	"strconv"
	"strings"
)

type PassResult int

const (
	FAIL PassResult = iota
	PASS
)

type TestResult struct {
	Test
	RunID int `json:"runid"`
}

func NewTestResult(testid, date string, result, runid int) *TestResult {
	tr := new(TestResult)
	tr.Pattern = testid
	tr.Date = date
	tr.Result = result
	tr.RunID = runid
	return tr
}

func (tr *TestResult) GetTest() *Test {
	return &tr.Test
}

func TestResultFromTester(tstr Tester) *TestResult {
	t := new(TestResult)
	t.Pattern = tstr.GetTest().Pattern
	t.Date = tstr.GetTest().Date
	t.Result = tstr.GetTest().Result
	t.RunID = 0
	return t
}

func ParseTestResults(tests []string) []*TestResult {
	trSlice := make([]*TestResult, len(tests))
	for i, test := range tests {
		tr := CreateTestResult(test)
		trSlice[i] = tr
	}
	return trSlice
}

func CreateTestResult(test string) *TestResult {
	split := strings.Split(test, ":")
	i, err := strconv.Atoi(split[2])
	HandleError(err)
	tr := NewTestResult(split[0], split[1], i, 0)
	return tr
}
