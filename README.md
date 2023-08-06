## Hexagonal Architecture
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
├── migrations
│   └── _init.sql
└── pkg
    └── server
        └── server.go
```

## Models
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