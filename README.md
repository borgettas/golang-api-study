### Run Project
``` bash
docker compose up --build [-d]
```

### Health API
``` bash
make health
```

### Sample Test
``` bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Armbrusteri","phone":123456789, "birthdate": "2000-01-01"}' http://localhost:8080/putter
```