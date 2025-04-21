SELECT * FROM users
WHERE id IN (
  SELECT sender_id FROM transactions
  WHERE note LIKE '%invoice%' OR note LIKE '%tax%'
  AND is_admin = 0
);
