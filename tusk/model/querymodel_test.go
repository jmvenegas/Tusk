package model

import (
	"testing"
)

func Test_QueryModel_TableFromQuery_0_Pass(t *testing.T) {
	cases := []struct {
		mysqlQuery, result string
	}{
		{
			"SELECT * FROM Results WHERE date = 2015-07-23",
			"Results",
		},
		{
			"SELECT * FROM MyTable WHERE result = 1",
			"MyTable",
		},
		{
			"SELECT * FROM Results",
			"Results",
		},
	}
	for _, c := range cases {
		returnStr := TableFromQuery(c.mysqlQuery)
		if returnStr != c.result {
			t.Errorf("TableFromQuery returned %s, expected %s: TestCase - 0", returnStr, c.result)
		}
	}
}
