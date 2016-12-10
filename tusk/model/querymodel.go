package model

import (
	"strings"

	"github.com/jmvenegas/tusk/parsing"
)

type Query struct {
	Test
	QueryType string `json:"querytype"`
}

func NewQuery(q, t, p, d string, r int) *Query {
	query := new(Query)
	query.QueryType = q
	query.Table = t
	query.Pattern = p
	query.Date = d
	query.Result = r
	return query
}

func (q *Query) GetTest() *Test {
	return &q.Test
}

func TableFromQuery(q string) string {
	if strings.Contains(q, "WHERE") {
		return parsing.WordBetweenStrings(q, "FROM ", " WHERE")
	}
	return q[strings.Index(q, "FROM ")+len("FROM "):]
}
