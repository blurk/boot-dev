INSERT INTO banks(id, name, routing_number)
VALUES (1, 'Central Savings', 123456789);

INSERT INTO banks(id, name, routing_number)
VALUES (2, 'Bank of Boots', 987654321);

INSERT INTO banks(id, name, routing_number)
VALUES (3, 'Metro Trust Bank', 456789123);

INSERT INTO users_banks(user_id, bank_id)
VALUES (1, 1);

INSERT INTO users_banks(user_id, bank_id)
VALUES (1, 2);

INSERT INTO users_banks(user_id, bank_id)
VALUES (2, 2);

INSERT INTO users_banks(user_id, bank_id)
VALUES (2, 3);

INSERT INTO users_banks(user_id, bank_id)
VALUES (3, 3);

INSERT INTO users_banks(user_id, bank_id)
VALUES (4, 3);

SELECT * FROM banks
WHERE id IN (
  SELECT bank_id FROM users_banks
);
