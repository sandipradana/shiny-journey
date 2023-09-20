## https://shinyjourney.fly.dev/


### auth / register
```
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Admin Name",
    "email": "admin@example.com",
    "password": "password123"
  }'
```

### auth / login
```
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

### employee / create
```
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "position": "Software Engineer"
  }'
```

### employee /  get
```
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### employee / update
```
curl -X PUT http://localhost:8080/employees/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated John Doe",
    "email": "updated_john@example.com"
  }'
```

### employee / delete
```
curl -X DELETE http://localhost:8080/employees/1
```

### attendance / create
```
curl -X POST http://localhost:8080/attendances \
  -H "Content-Type: application/json" \
  -d '{
    "employee_id": 1,
    "attendance_date": "2023-09-21T00:00:00Z",
    "status": "present"
  }'
```

### attendance / get
```
curl http://localhost:8080/attendances/1
```

### attendance / update
```
curl -X PUT http://localhost:8080/attendances/1 \
  -H "Content-Type: application/json" \
  -d '{
    "employee_id": 2,
    "attendance_date": "2023-09-21T12:34:56Z",
    "status": "absent"
  }'
```

### attendace / delete
```
curl -X DELETE http://localhost:8080/attendances/1
```