apply:
	cat feed.sql | psql postgres://staging:staging@localhost/staging?sslmode=disable