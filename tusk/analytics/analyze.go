package analytics

import (
	"github.com/jmvenegas/tusk/model"
)

func GetPassingTests(testsIn []*model.TestResult) []*model.TestResult {
	return getTestsWithResult(testsIn, model.PASS)
}

func GetFailingTests(testsIn []*model.TestResult) []*model.TestResult {
	return getTestsWithResult(testsIn, model.FAIL)
}

func getTestsWithResult(testsIn []*model.TestResult, result model.PassResult) []*model.TestResult {
	testsOut := make([]*model.TestResult, 0)
	for _, t := range testsIn {
		if t.Result == int(result) {
			testsOut = append(testsOut, t)
		}
	}
	return testsOut
}
