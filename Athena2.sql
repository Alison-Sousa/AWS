-- Criação de tabela fictícia para registros de vendas
CREATE EXTERNAL TABLE IF NOT EXISTS sales (
    transaction_id INT,
    product_name STRING,
    category STRING,
    price DECIMAL(10, 2),
    quantity INT,
    transaction_date TIMESTAMP
)
ROW FORMAT DELIMITED
FIELDS TERMINATED BY ','
LOCATION 's3://my-bucket/sales/';

-- Inserção de dados fictícios na tabela de vendas
INSERT INTO sales VALUES
    (1, 'Product A', 'Electronics', 500.00, 2, '2023-01-15 10:30:00'),
    (2, 'Product B', 'Clothing', 35.00, 3, '2023-01-15 11:45:00'),
    (3, 'Product C', 'Electronics', 1200.00, 1, '2023-01-16 09:15:00'),
    (4, 'Product D', 'Books', 20.00, 5, '2023-01-16 14:20:00'),
    (5, 'Product E', 'Home & Kitchen', 300.00, 2, '2023-01-17 13:00:00');

-- Consulta para selecionar todas as vendas de eletrônicos
SELECT *
FROM sales
WHERE category = 'Electronics';

-- Consulta para calcular o total de vendas por categoria
SELECT category, SUM(price * quantity) AS total_sales
FROM sales
GROUP BY category
ORDER BY total_sales DESC;

-- Consulta para obter o número de transações por dia
SELECT DATE(transaction_date) AS transaction_day, COUNT(*) AS transaction_count
FROM sales
GROUP BY DATE(transaction_date)
ORDER BY transaction_day;
