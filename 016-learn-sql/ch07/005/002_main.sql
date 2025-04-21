SELECT user_id, SUM(amount) as balance FROM transactions
WHERE was_successful = true
GROUP BY user_id;