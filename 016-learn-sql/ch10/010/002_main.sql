SELECT
  users.name,
  users.username,
  COUNT(support_tickets.id) as support_ticket_count
FROM
  users
INNER JOIN
  support_tickets ON users.id = support_tickets.user_id
WHERE
  issue_type IS NOT 'Account Access'
GROUP BY
  users.id
HAVING
  support_ticket_count > 1
ORDER BY
  support_ticket_count DESC;