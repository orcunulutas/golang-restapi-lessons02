[![text](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/orcunulutas/)

[![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/bukotsunikki.svg?style=social&label=Follow%20%40oulutas)](https://twitter.com/oulutas)


## Hexagonal Architecture
database : postgres:latest

```golang
my-app
├── cmd
│   └── main.go
├── internal
│   ├── adapter
│   │   ├── database
│   │   │   └── database.go
│   │   └── http
│   │       ├── handler.go
│   │       └── router.go
│   ├── domain
│   │   ├── model
│   │   │  ├── task_model.go
│   │   │  └── user_model.go
│   │   ├── task_repository.go
│   │   ├── user_repository.go
│   └── usecase
│       ├── task_usecase.go
│       └── user_usecase.go
└── pkg
    └── server
        └── server.go
```

## GO Modules

```ssh
go mod init go-restapi
go get -u github.com/go-chi/chi/v5
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/stretchr/testify/assert
```

## Models

Task table;
```golang
type Task struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
```

```sql
CREATE TABLE public.tasks (
    id integer NOT NULL,
    title character varying(512),
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
```

User Table;
```golang
type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```
```sql
CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
```
