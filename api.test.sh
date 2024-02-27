# Test for getting list of users
curl http://localhost:8080/users/

# Test for getting specific user
curl http://localhost:8080/users/4

# Test for creating a user
curl -H "Content-Type: application/json" \
     -X POST http://localhost:8080/users/Mitrajeet

# Test for updating a user
curl -H "Content-Type: application/json" \
     -d '{"name": "Mitrajeet 2"}' \
     -X PUT http://localhost:8080/users/4

# Test for deleting a user
curl -H "Content-Type: application/json" \
     -X PUT http://localhost:8080/users/4
