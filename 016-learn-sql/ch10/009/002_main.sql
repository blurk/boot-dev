SELECT users.id, users.name, users.age, users.username, countries.name AS country_name, SUM(transactions.amount) AS balance
FROM users
INNER JOIN countries
ON users.country_code = countries.country_code
LEFT JOIN transactions
ON users.id = transactions.user_id AND transactions.was_successful = 1
WHERE users.id = 3;