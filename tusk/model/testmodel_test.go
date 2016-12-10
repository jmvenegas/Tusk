package model

import (
	"reflect"
	"testing"
)

func Test_TestModel_TestResultFromTester_0_Pass(t *testing.T) {
	cases := []struct {
		testResult *TestResult
		tester     Tester
	}{
		{
			NewTestResult("testname1", "2015", 1, 0),
			NewQuery("queryall", "Results", "testname1", "2015", 1),
		},
	}
	for _, c := range cases {
		retResult := TestResultFromTester(c.tester)
		if reflect.DeepEqual(c.testResult, retResult) != true {
			t.Errorf("TestResultFromTester failed to return expected TestResult\n%s != %s\n: TestCase - 0", c.testResult, retResult)
		}
	}
}

func Test_TestModel_CreateTestResult_1_Pass(t *testing.T) {
	cases := []struct {
		testResult *TestResult
		testStr    string
	}{
		{
			NewTestResult("testname1", "2015", 1, 0),
			"testname1:2015:1",
		},
	}
	for _, c := range cases {
		retResult := CreateTestResult(c.testStr)
		if reflect.DeepEqual(c.testResult, retResult) != true {
			t.Errorf("CreateTestResult failed to return expected TestResult\n%s != %s\n: TestCase - 1", c.testResult, retResult)
		}
	}
}
