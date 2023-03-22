# Periodic Task API

---

### Start app command via docker

```sh
make docker.app.start
``` 

*This command will start the app with *port* and *address* defined as CLI arguments when CMD will be executed
inside `deployment/production/Dockerfile`

### Start app command via cli

```sh
go run ./cmd/main/main.go -p 8080 -addr 127.0.0.1
```

---

### Endpoints

<details>
 <summary><code>GET</code> <code><b>/ptlist/{period}{tz}{t1}{t2}</b></code> <code>(Returns matching timestamps of a periodic task)</code></summary>

##### Parameters

|   name   |    type    | data type |       description       |
|:--------:|:----------:|:---------:|:-----------------------:|
| `period` | *required* | *string*  | Period of periodic task |
|   `tz`   | *required* | *string*  |   Timestamps timezone   |
|   `t1`   | *required* | *string*  |     Start timestamp     |
|   `t2`   | *required* | *string*  |      End timestamp      |

**period** parameter allowed values

- 1h
- 1d
- 1mo
- 1y

##### Responses

| http code |           content-type            |                  response                  | description                 |
|:---------:|:---------------------------------:|:------------------------------------------:|:----------------------------|
|   `200`   | `application/json; charset=utf-8` | `["20191231T220000Z", "20201231T220000Z"]` | Valid request and response  |
|   `400`   | `application/json; charset=utf-8` |  `{"status": "string", "desc": "string"}`  | Invalid request's url param |
|   `500`   | `application/json; charset=utf-8` |  `{"status": "string", "desc": "string"}`  | Internal app error          |

##### Example

```sh
http://127.0.0.1:8080/ptlist?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z
```

---
</details>

### Makefile Commands

| Command                               | Usage                                                      |
|---------------------------------------|------------------------------------------------------------|
| docker.app.start                      | `Start all services`                                       |
| docker.test.unit                      | `Run unit tests via docker`                                |
| docker.test.all                       | `Run both unit and integration tests via docker`           |
| docker.test.all.coverage.withView     | `Run both unit and integration tests via docker with view` |
| docker.mock.generate FILE={FILE_PATH} | `Generate mock for a specific file via docker`             |

* In order to execute makefile commands type **make** plus a command from the table above

  make {command}

---

### Notes

- *config.yaml* has been committed and pushed for the assessment needs. Config and env files should **always** be
  untracked and never been committed.

---