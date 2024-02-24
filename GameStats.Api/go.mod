module GameStats.Api

go 1.21.3

require (
	github.com/go-chi/chi v1.5.5
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

require GameStats.Domain v0.0.0
require GameStats.Infrastructure v0.0.0

replace GameStats.Domain => ../GameStats.Domain

replace GameStats.Infrastructure => ../GameStats.Infrastructure
