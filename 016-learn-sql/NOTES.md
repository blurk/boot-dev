While most relational databases are fairly similar, NoSQL databases tend to be fairly unique and are used for more niche purposes.

Some of the main differences between a SQL and NoSQL databases are:

- NoSQL databases are usually non-relational, SQL databases are usually relational (we'll talk more about what this means later).
- SQL databases usually have a defined schema, NoSQL databases usually have dynamic schema.
- SQL databases are table-based, NoSQL databases have a variety of different storage methods, such as document, key-value, graph, wide-column, and more.

Types of NoSQL Databases
- Document Database
- Key-Value Store
- Wide-Column
- Graph

A few of the most popular NoSQL databases are:
- MongoDB
- Cassandra
- CouchDB
- DynamoDB
- ElasticSearch

## Intro to Migrations

A database migration is **a set of changes** to a relational database. In fact, the ALTER TABLE statements we did in the last exercise were examples of migrations!

Migrations are helpful when transitioning from one state to another, fixing mistakes, or adapting a database to changes.

Good migrations are small, incremental and ideally reversible changes to a database. As you can imagine, when working with large databases, making changes can be scary! We have to be careful when writing database migrations so that we don't break any systems that depend on the old database schema.

When writing reversible migrations, we use the terms "up" and "down" migrations. An "up" migration is simply the set of changes you want to make, like altering/removing/adding/editing a table in some way. A "down" migration includes the changes that would revert any of the "up" migration's changes.

## Constrains

In SQL, a cell with a NULL value indicates that the value is missing. A NULL value is very different from a zero value.

A constraint is a rule we create on a database that enforces some specific behavior. For example, setting a NOT NULL constraint on a column ensures that the column will not accept NULL values.

If we try to insert a NULL value into a column with the NOT NULL constraint, the insert will fail with an error message. Constraints are extremely useful when we need to ensure that certain kinds of data exist within our database.

## Keys

A key defines and protects relationships between tables. A primary key is a special column that uniquely identifies records within a table. Each table can have one, and only one primary key.

Foreign keys are what makes relational databases relational! Foreign keys define the relationships between tables. Simply put, a FOREIGN KEY is a field in one table that references another table's PRIMARY KEY.

## There Is No Perfect Way to Architect a Database Schema

When designing a database schema there typically isn't a "correct" solution. We do our best to choose a sane set of tables, fields, constraints, etc that will accomplish our project's goals. Like many things in programming, different schema designs come with different tradeoffs.

## Relational Databases

A relational database is a type of database that stores data so that it can be easily related to other data.

For example, a user can have many tweets. There's a relationship between a user and their tweet.

In a relational database:

1. Data is typically represented in "tables".
2. Each table has "columns" or "fields" that hold attributes related to the record.
3. Each row or entry in the table is called a record.
4. Typically, each record has a unique Id called the primary key.

![example](image.png)

## Relational vs. Non-Relational Databases

The big difference between relational and non-relational databases is that non-relational databases nest their data. Instead of keeping records on separate tables, they store records within other records.

To over-simplify it, you can think of non-relational databases as giant JSON blobs. If a user can have multiple courses, you might just add all the courses to the user record.

This often results in duplicate data within the database. That's obviously less than ideal, but it does have some benefits.
