SELECT country_code, ROUND(AVG(age), 0) AS average_age FROM users
GROUP BY country_code;
