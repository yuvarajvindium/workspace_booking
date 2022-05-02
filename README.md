
## Create Users Table 

```sql
CREATE TABLE public.users
(
    id uuid NOT NULL,
    email character varying(254) NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    PRIMARY KEY (id)
);
```

Luve run go
`nodemon --exec go run main.go --signal SIGTERM`