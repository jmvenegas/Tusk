package model

import (
	"reflect"
	"testing"
)

func Test_DatabaseManager_GetQueryFunction_0_Pass(t *testing.T) {
	cases := []struct {
		origin   *DatabaseManager
		queryStr string
	}{
		{
			NewDatabaseManager("user", "pass", ":8585", "mysql", "Results"),
			"QueryAll",
		},
	}
	for _, c := range cases {
		refValue1 := reflect.ValueOf(c.origin.QueryAll).String()
		refValue2 := reflect.ValueOf(c.origin.GetQueryFunction(c.queryStr)).String()
		if refValue1 != refValue2 {
			t.Errorf("GetQueryFunction:%s and %s differ: TestCase - 0", refValue1, refValue2)
		}
	}
}
