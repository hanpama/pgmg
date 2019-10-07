// Code Generated by github.com/hanpama/pgmg
package queries

import "time"

const SQL = `
SELECT orders.*
FROM orders
  INNER JOIN customers ON customers.country = orders.ship_country
WHERE customers.country = $1
`

func Query(args Args) query {
	return query{args}
}

type Args struct {
	a1 *string
}

type query struct{ Args }

func (q query) SQL() string { return SQL }
func (q query) Args() []interface{} {
	return []interface{}{q.a1}
}

type Record struct {
	OrderID        *int16     `json:"order_id"`
	CustomerID     *string    `json:"customer_id"`
	EmployeeID     *int16     `json:"employee_id"`
	OrderDate      *time.Time `json:"order_date"`
	RequiredDate   *time.Time `json:"required_date"`
	ShippedDate    *time.Time `json:"shipped_date"`
	ShipVia        *int16     `json:"ship_via"`
	Freight        *float32   `json:"freight"`
	ShipName       *string    `json:"ship_name"`
	ShipAddress    *string    `json:"ship_address"`
	ShipCity       *string    `json:"ship_city"`
	ShipRegion     *string    `json:"ship_region"`
	ShipPostalCode *string    `json:"ship_postal_code"`
	ShipCountry    *string    `json:"ship_country"`
}

func (r *Record) Receive() []interface{} {
	return []interface{}{
		&r.OrderID,
		&r.CustomerID,
		&r.EmployeeID,
		&r.OrderDate,
		&r.RequiredDate,
		&r.ShippedDate,
		&r.ShipVia,
		&r.Freight,
		&r.ShipName,
		&r.ShipAddress,
		&r.ShipCity,
		&r.ShipRegion,
		&r.ShipPostalCode,
		&r.ShipCountry,
	}
}

type Recordset []Record

func (rs *Recordset) ReceiveNext() []interface{} {
	*rs = append(*rs, Record{})
	return (*rs)[len(*rs)-1].Receive()
}