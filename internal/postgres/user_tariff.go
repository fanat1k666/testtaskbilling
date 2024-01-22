package postgres

import (
	"TestTask/internal/repository"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(host string, port int, user string, password string, dbname string) (*Postgres, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			host, port, user, password, dbname),
	)
	if err != nil {
		return nil, fmt.Errorf("can't open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping db: %w", err)
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) ShowTariff(userId int) ([]repository.ShowUserTariff, error) {
	rows, err := p.db.Query(`SELECT u.user_id, name FROM tariffs
	INNER JOIN user_tariffs ON tariffs.id = user_tariffs.tariffs_id
	INNER JOIN public."user" u on user_tariffs.user_id = u.id
	WHERE u.user_id = $1`, userId)
	if err != nil {
		return nil, fmt.Errorf("can't found user: %w", err)
	}
	defer rows.Close()
	var tariffs []repository.ShowUserTariff
	for rows.Next() {
		s := repository.ShowUserTariff{}
		err = rows.Scan(&s.UserId, &s.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tariffs = append(tariffs, s)
	}
	return tariffs, nil
}

func (p *Postgres) CreateTariff(name string, price int) error {
	_, err := p.db.Exec(`INSERT INTO tariffs (name,price) VALUES ($1,$2)`, name, price)
	if err != nil {
		return fmt.Errorf("can't create tariff: %w", err)
	}
	return nil
}

func (p *Postgres) UpdateTariff(name string, price int) error {
	_, err := p.db.Exec(`UPDATE tariffs SET price=$2 WHERE name=$1`, name, price)
	if err != nil {
		return fmt.Errorf("can't update tariff: %w", err)
	}
	return nil
}
