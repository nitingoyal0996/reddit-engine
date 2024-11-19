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

Login User with -

```bash

# request
❯ curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "username": "johndoe",
    "password": "secretPass123"
}'

# response
{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJSZWRkaXQiLCJleHAiOjE3MzIwNTc1NDksImlhdCI6MTczMjA1NjY0OSwidXNlcl9pZCI6MiwidXNlcm5hbWUiOiJqb2huZG9lIn0.Mu2akAa3Q0b89rl1qkYs99dMCwpnMuZVhieeaigKHAI","Error":""}

```
