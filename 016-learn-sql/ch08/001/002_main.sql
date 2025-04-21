SELECT *
FROM transactions
WHERE user_id IN (
  SELECT user_id
  FROM users
  WHERE name = 'David'
);