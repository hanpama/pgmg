SELECT orders.*
FROM orders
  INNER JOIN customers ON customers.country = orders.ship_country
WHERE customers.country = $1
