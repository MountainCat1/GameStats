module GameStats.Api

go 1.21

require (
	github.com/go-chi/chi v1.5.5
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-openapi/jsonpointer v0.20.2 // indirect
	github.com/go-openapi/jsonreference v0.20.4 // indirect
	github.com/go-openapi/spec v0.20.14 // indirect
	github.com/go-openapi/swag v0.22.9 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.4 // indirect
	gorm.io/driver/sqlite v1.5.5 // indirect
	gorm.io/gorm v1.25.7 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

require GameStats.Domain v0.0.0

require (
	GameStats.Infrastructure v0.0.0
	github.com/swaggo/http-swagger v1.3.4
	github.com/swaggo/swag v1.16.3
)

replace GameStats.Domain => ../GameStats.Domain

replace GameStats.Infrastructure => ../GameStats.Infrastructure
