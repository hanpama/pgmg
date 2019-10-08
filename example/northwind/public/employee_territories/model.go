// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package employee_territories

import (
	"encoding/json"
)

type Record struct {
	EmployeeID  int16  `json:"employee_id"`
	TerritoryID string `json:"territory_id"`
}

type PkEmployeeTerritories struct {
	EmployeeID  int16  `json:"employee_id"`
	TerritoryID string `json:"territory_id"`
}

type EmployeeID int16
type TerritoryID string

func (r *Record) Receive() []interface{} {
	return []interface{}{
		&r.EmployeeID,
		&r.TerritoryID,
	}
}

type Recordset []Record

func (rs *Recordset) ReceiveNext() []interface{} {
	*rs = append(*rs, Record{})
	return (*rs)[len(*rs)-1].Receive()
}

type Values []attribute

func InputValues(
	employeeID EmployeeID,
	territoryID TerritoryID,
	attrs ...attribute,
) Values {
	return append(Values{
		employeeID,
		territoryID,
	}, attrs...)
}
func (vs Values) ApplyTo(r *Record) {
	for _, v := range vs {
		v.ApplyTo(r)
	}
}

func (vs Values) MarshalJSON() (b []byte, err error) {
	r := make(map[string]interface{})
	for _, v := range vs {
		r[v.Column()] = v.Value()
	}
	return json.Marshal(r)
}

type attribute interface {
	ApplyTo(*Record)
	Column() string
	Value() interface{}
}

func (v EmployeeID) ApplyTo(r *Record)   { r.EmployeeID = (int16)(v) }
func (v EmployeeID) Column() string      { return "employee_id" }
func (v EmployeeID) Value() interface{}  { return (int16)(v) }
func (v TerritoryID) ApplyTo(r *Record)  { r.TerritoryID = (string)(v) }
func (v TerritoryID) Column() string     { return "territory_id" }
func (v TerritoryID) Value() interface{} { return (string)(v) }

func mustMarshalJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
