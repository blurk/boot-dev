SELECT sender_id, SUM(amount) as balance
FROM transactions
WHERE was_successful = true
AND sender_id IS NOT NULL
AND note LIKE '%lunch%'
GROUP BY sender_id
HAVING balance > 20
ORDER BY balance;