package utilities

import (
	"strconv"
	"strings"
)

type NullableFloat struct {
	Valid bool
	Value float64
}

func (nf *NullableFloat) UnmarshalCSV(s string) error {
	if strings.TrimSpace(s) == "" {
		nf.Valid = false
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	nf.Valid = true
	nf.Value = v
	return nil
}
