SELECT COUNT(*) AS junior_count FROM users
WHERE (country_code = 'US' or country_code = 'CA') AND age < 18;