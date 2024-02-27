module gamestats-intrastructure

go 1.21

require (
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gorm.io/driver/mysql v1.5.4
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.7
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
)

require gamestats-domain v0.0.0

replace gamestats-domain => ./../gamestats-domain