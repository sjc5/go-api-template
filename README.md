# go-api-template

## Getting Started

1. Clone the repo
2. Change the name of the module in go.mod to whatever you want
3. Run `go mod tidy`
4. Run `make run` or `go run ./cmd/app` in your terminal
5. Visit `http://localhost:8080` in your browser

A few things to try:

- At `/` you should see the heartbeat route. It just returns a `.` and a 200 status code.
- At `/public` you should get a 200 response. This is a public route.
- At `/protected` you have a 50/50 chance of getting a 401 or a 200 (there is a faux session getter that calls a random number generator). This is a protected route.
- If you rapidly refresh your browser at any route, you should get a 429 from the rate limiter middleware.

## Some Notes

### Conventions

- Vars ending in "Size" means the size in bytes

### Auth

This template sets up a tiny bit of scaffolding show how you might interact with sessions using middleware, but it doesn't set up an auth system at all. Maybe you want cookies, maybe you want passkeys, maybe you want JWTs... who knows. So that's your responsibility. If you go with cookies, make sure you handle setting up CSRF protection and cookie signing.

### Database

This template doesn't actually do anything with the database, but it just shows how to set it up. The instance is a global variable available at `global.DB`.

#### Migrations

To auto-migrate the database, run `make migrate` or `go run ./cmd/migrate`. This will create the tables and columns in the database. When you add a new model, you'll need to pass it to the auto-migrator in the migrate script in `cmd/migrate/main.go`.

#### PlanetScale

If you're using PlanetScale:

1. Replace `gorm.io/driver/sqlite` with `gorm.io/driver/mysql`.

2. Do this for your gorm.Open call:

```go
global.DB, err = gorm.Open(mysql.Open(env.DSN), &gorm.Config{
  DisableForeignKeyConstraintWhenMigrating: true,
})
```

3. Use this DSN format: `ps_conn_string?tls=true&interpolateParams=true`

### Validation

This template doesn't do any validation, but it does show how to set up a global validator instance. Additionally, there are a couple utilities in the utils package to help with validating user input.

## License

MIT
