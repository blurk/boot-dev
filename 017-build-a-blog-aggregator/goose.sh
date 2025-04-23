cd  sql/schema
goose postgres postgres://postgres:sinh@localhost:5432/gator down
goose postgres postgres://postgres:sinh@localhost:5432/gator up
cd ../..