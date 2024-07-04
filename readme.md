# Endpoint:

- Login (/login)
- Create Customer (/create)

## Usage

1.  Copy .env.example to .env and set the environment variables:

    `cp .env.example .env`

2.  Run your application using the command in the terminal:

    `docker-compose up -d --build`

3.  Hit url login to get JWT token, example:

    `curl  -X GET 'http://127.0.0.1:3000/login' --header 'Accept: */*' --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'`

4.  Create a new customer with authorization, example:

    `curl  -X POST 'http://127.0.0.1:3000/create' --header 'Accept: */*' --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjAyMzQwNTAsInVzZXIiOiJudXJoYWtpa2kifQ.SuhDU0gJP5A9eH7QNDDzDQU0BWvEBaBd68_Q20vaLrE' --header 'Content-Type: application/json' --data-raw '{"parent_id": "99682a9d-83bc-4766-bd3b-4ec236f1351b"}'`
