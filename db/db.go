/* db — это фактическая реализация взаимодействия с базой данных. Здесь конструируются и исполняются операторы SQL.
Этот пакет также импортирует model, он должен будет создать эти структуры из данных базы данных.
В первую очередь, db должен предоставить функцию InitDB, которая установит соединение с базой данных,
а также создаст необходимые таблицы и подготовит SQL запросы.
*/

package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gowebapp/model"
)

type Config struct {
	ConnectString string
}

// InitDb() создает экземпляр pgDb, который является Postgres-реализацией нашего интерфейса model.db
func InitDb(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nil {
			return nil, err
		}
		if err := p.createTablesIfNotExist(); err != nil {
			return nil, err
		}
		if err := p.prepareSqlStatements(); err != nil {
			return nil, err
		}

		return p, nil
	}
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectPeople *sqlx.Stmt
}

// createTablesIfNotExist для создания таблиц
func (p *pgDb) createTablesIfNotExist() error {
	createSql := `
		CREATE TABLE IF NOT EXISTS people (
			id SERIAL NOT NULL PRIMARY KEY,
			first TEXT NOT NULL,
			last TEXT NOT NULL);
		INSERT INTO people (first, last) VALUES('John', 'Doe');
	`
	if rows, err := p.dbConn.Query(createSql); err != nil {
		return err
	} else {
		rows.Close()
	}
	return nil
}

// prepareSqlStatements для подготовки запросов
func (p *pgDb) prepareSqlStatements() (err error) {
	if p.sqlSelectPeople, err = p.dbConn.Preparex(
		"SELECT id, first, last FROM people",
	); err != nil {
		return err
	}

	return nil
}

// SelectPeople - метод, реализующий интерфейс
func (p *pgDb) SelectPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	if err := p.sqlSelectPeople.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}
