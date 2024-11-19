# Reddit-engine

build: `go build .`
run: `./reddit-clone`

Register user with -

```bash
# request
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "username": "johndoe",
    "email": "john.doe@example.com",
    "password": "secretPass123"
}'

# response
{"ID":2,"Username":"johndoe","Error":""}
```
