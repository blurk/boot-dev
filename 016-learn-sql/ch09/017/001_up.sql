CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  bank_id INTEGER,
  bank_name TEXT,
  bank_routing_number INTEGER,
  is_admin BOOLEAN
);

INSERT INTO users(name, age, username, password, is_admin, bank_id, bank_name, bank_routing_number)
VALUES ('David', 34, 'david.lang', 'secure1234', false, 1, 'Central Savings', 123456789);

INSERT INTO users(name, age, username, password, is_admin, bank_id, bank_name, bank_routing_number)
VALUES ('Sam', 12, 'sam-show', 'nasjds134', false, 2, 'Bank of Boots', 987654321);

INSERT INTO users(name, age, username, password, is_admin, bank_id, bank_name, bank_routing_number)
VALUES ('Lane', 19, 'wagslane', '2jk3bAkm', false, 3, 'Metro Trust Bank', 456789123);

INSERT INTO users(name, age, username, password, is_admin, bank_id, bank_name, bank_routing_number)
VALUES ('Allan', 27, 'allan.jules', '243nldn', false, 2, 'Bank of Boots', 987654321);
