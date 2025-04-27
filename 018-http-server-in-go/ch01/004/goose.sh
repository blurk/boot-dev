echo "down start"
cd  sql/schema
goose postgres postgres://postgres:sinh@localhost:5432/chirpy down
echo "down end"
echo "up start"
goose postgres postgres://postgres:sinh@localhost:5432/chirpy up
cd ../..
echo "up end"