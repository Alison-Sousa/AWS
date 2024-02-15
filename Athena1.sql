-- Criação da tabela fictícia de funcionários
CREATE EXTERNAL TABLE IF NOT EXISTS employees (
    id INT,
    name STRING,
    position STRING,
    salary DECIMAL(10, 2)
)
ROW FORMAT DELIMITED
FIELDS TERMINATED BY ','
LOCATION 's3://my-bucket/employees/';

-- Inserção de dados fictícios na tabela de funcionários
INSERT INTO employees VALUES
    (1, 'John', 'Developer', 6000.00),
    (2, 'Alice', 'Manager', 8000.00),
    (3, 'Bob', 'Designer', 5500.00);

-- Consulta para selecionar todos os funcionários
SELECT * FROM employees;

-- Consulta para selecionar funcionários com salário superior a $5000
SELECT * FROM employees WHERE salary > 5000;

-- Consulta para contar o número de funcionários por cargo
SELECT position, COUNT(*) AS employee_count
FROM employees
GROUP BY position;
