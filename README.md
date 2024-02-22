# go-api-template

## Getting Started

1. Clone the repo
2. Change your module name in `go.mod`
3. Run `go mod tidy`
4. Run `make run` or `go run ./cmd/app`
5. Visit `http://localhost:8080`

A few things to try:

- At `/`, the response should be `.` with a `200` status code. This is the healthcheck route set up by the heartbeat middleware.
- At `/public`, the response should be `Welcome` with a `200` status code. This is a public route (no session required).
- At `/protected`, you have a 50/50 chance of getting an `Unauthorized` response with a `401` status code, or a response showing you a fake User ID with a `200` status code. This is a protected route (session required). There's just a faux session getter using a random number generator here – replace with your actual logic.
- If you rapidly refresh your browser at any route, you should get a `Too Many Requests` response with a `429` status code from the rate limiter middleware. This is set ultra conservatively at 1 request per second just to make it easy to trigger for demo purposes.

## Some Notes

### Conventions

- Vars ending in "Size" means the size in bytes

### Auth

This template sets up a tiny bit of scaffolding show how you might interact with sessions using middleware, but it doesn't set up an auth system at all. Maybe you want cookies, maybe you want passkeys, maybe you want JWTs, who knows. So that's your responsibility. If you go with cookies, make sure you handle setting up CSRF protection (and perhaps cookie signing).

### Database

This template doesn't actually do anything with the database, but it just shows how to set one up with GORM. The instance is a global variable at `global.DB`.

#### Migrations

To auto-migrate the database, run `make migrate` or `go run ./cmd/migrate`. This will create the tables and columns in the database. When you add a new model, don't forget to add it to the migrator script at `cmd/migrate/main.go`.

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

This template doesn't do any validation, but it does show how to set up a global validator instance (available at `global.Validate`). Additionally, there are a couple utilities in the utils package to help with validating user input (`UnmarshalAndValidateFromRequest` and `UnmarshalAndValidateFromString`).

## License

MIT
