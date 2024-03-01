# go-api-template

## Getting Started

1. Clone the repo
2. Change your module name in `go.mod`
3. Run `go mod tidy`
4. Run `make dev`
5. Visit `http://localhost:8080`

A few things to try:

- At `/`, the response should be `.` with a `200` status code. This is the healthcheck route set up by the heartbeat middleware.
- At `/public`, the response should be `Welcome` with a `200` status code. This is a public route (no session required).
- At `/protected`, you have a 50/50 chance of getting an `Unauthorized` response with a `401` status code, or a response showing you a fake User ID with a `200` status code. This is a protected route (session required). There's just a faux session getter using a random number generator here – replace with your actual logic.
- If you rapidly refresh your browser at any route, you should get a `Too Many Requests` response with a `429` status code from the rate limiter middleware. For demo purposes, this is set at 1 request per second to make it easy to trigger.

## Some Notes

### Dev tooling

This template uses [Kiruna](https://github.com/sjc5/kiruna) for dev server reloads. If you're building a full-stack app, Kiruna can do a lot more (like hot CSS reloading, browser refreshes, static asset hashing, etc.), but in this case we are just using it as an alternative to [Air](https://github.com/cosmtrek/air). The nice thing about Kiruna for this use case is that it doesn't require you to install any extra tooling on your machine, while still having a very light footprint in the repo.

### Conventions

- Vars ending in "Size" means the size in bytes

### Auth

This template sets up a bit of scaffolding to show how you might interact with sessions using context, but it doesn't set up an actual authentication system at all. So that's your responsibility.

### Database

This template doesn't actually do anything interesting with the database, but it does show how to set up a SQLite database using GORM. The GORM instance is a global variable at `platform.DB`.

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

This template doesn't do any validation, but it does show how to set up a global validator instance (available at `platform.Validate`). Additionally, there are a couple nice utility methods for simultaneously marshaling and validating JSON that are nice to know about: `platform.Validate.UnmarshalFromRequest()` and `platform.Validate.UnmarshalFromString()`. This uses `go-playground/validator` under the hood, so use the tags and syntax from that library.

## License

MIT
