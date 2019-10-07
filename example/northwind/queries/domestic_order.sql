SELECT
  orders.order_id,
  orders.order_date,
  orders.shipped_date,
  orders.ship_region,
  orders.ship_country
FROM orders
  INNER JOIN customers ON customers.country = orders.ship_country
WHERE customers.country = $1
ORDER BY orders.shipped_date DESC
LIMIT $2
