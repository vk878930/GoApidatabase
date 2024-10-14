/customer/
INSERT INTO customer (name, phone, address, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING customer_id;

SELECT * FROM customer;

SELECT * FROM customer WHERE customer_id=$1;

UPDATE customer SET name=$2, phone=$3, address=$4, updated_at=CURRENT_TIMESTAMP WHERE customer_id=$1;

DELETE FROM customer WHERE customer_id=$1;

/employee/
INSERT INTO employee (name, phone, address, created_at, updated_at) 
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
RETURNING employee_id;

SELECT * FROM employee;

SELECT * FROM employee WHERE employee_id=$1;

UPDATE customer SET name=$2, phone=$3, address=$4, updated_at=CURRENT_TIMESTAMP WHERE customer_id=$1;

DELETE FROM employee WHERE employee_id=$1;

/product/
INSERT INTO product(name, price, unit, created_at, updated_at)  
VALUES($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
RETURNING product_id

SELECT * FROM product

SELECT * FROM product WHERE product_id=$1;

UPDATE product SET name=$2, price=$3, unit=$4, updated_at=CURRENT_TIMESTAMP WHERE product_id=$1

DELETE FROM product WHERE product_id=$1;

/transactions & bill details/

INSERT INTO transaction(bill_date, entry_date, finish_date, employee_id, customer_id)
VALUES ($1, $2, $3, $4, $5) 
RETURNING transaction_id;

SELECT price FROM product WHERE product_id = $1;

INSERT INTO bill_details (transaction_id, product_id, product_price, qty)
VALUES ($1, $2, $3, $4);

SELECT 
t.transaction_id, 
t.bill_date, 
t.entry_date, 
t.finish_date, 
t.employee_id, 
t.customer_id,
COALESCE(SUM(b.product_price * b.qty), 0) AS total_bill
FROM transaction t
LEFT JOIN bill_details b ON t.transaction_id = b.transaction_id
GROUP BY t.transaction_id;


SELECT id, transaction_id, product_id, product_price, qty FROM bill_details WHERE transaction_id=$1

SELECT 
t.transaction_id, 
t.bill_date, 
t.entry_date, 
t.finish_date, 
t.employee_id, 
t.customer_id,
COALESCE(SUM(b.product_price * b.qty), 0) AS total_bill
FROM transaction t
LEFT JOIN bill_details b ON t.transaction_id = b.transaction_id
WHERE t.transaction_id = $1
GROUP BY t.transaction_id;

SELECT id, transaction_id, product_id, product_price, qty FROM bill_details WHERE transaction_id=$1