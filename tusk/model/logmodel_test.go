package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_LogEntry_ToJson_0_Pass(t *testing.T) {
	cases := []struct {
		origin LogEntry
	}{
		{
			LogEntry{
				Message:  "Fatal Error",
				Date:     time.Now(),
				LogLevel: FATAL,
			},
		},
	}
	for _, c := range cases {
		jsonBack := c.origin.ToJson()
		remadeStruct := LogEntry{}
		json.Unmarshal(jsonBack, &remadeStruct)
		if reflect.DeepEqual(remadeStruct, c.origin) != true {
			t.Errorf("Json struct did not match original struct: TestCase - 0")
		}
	}
}

func Test_LogEntry_ToString_1_Pass(t *testing.T) {
	timenow := time.Now()
	cases := []struct {
		origin       LogEntry
		originString string
	}{
		{
			LogEntry{
				Message:  "Fatal Error",
				Date:     timenow,
				LogLevel: FATAL,
			},
			fmt.Sprintf("[FATA]%s : Fatal Error", timenow.Format(time.RFC850)),
		},
	}
	for _, c := range cases {
		stringBack := c.origin.ToString()
		if stringBack != c.originString {
			t.Errorf("ToString returned an unmatching string representaiton: TestCase - 1")
		}
	}
}

func Test_LogModel_SetLogLevel_2_Pass(t *testing.T) {
	cases := []struct {
		origin LogModel
	}{
		{
			LogModel{
				FileName:     "testFile",
				FP:           nil,
				CurrentLevel: FATAL,
			},
		},
	}
	for _, c := range cases {
		c.origin.SetLogLevel(WARN)
		logLevel := c.origin.CurrentLevel
		if logLevel != WARN {
			t.Errorf("SetLogLevel failed. Set to level %d but was %d: TestCase - 2", WARN, logLevel)
		}
	}
}
