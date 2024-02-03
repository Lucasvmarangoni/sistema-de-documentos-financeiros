package repositories

import (
	"context"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Insert(user *entities.User, ctx context.Context) error
	FindByEmail(email string, ctx context.Context) (*entities.User, error)
	FindById(id pkg_entities.ID, ctx context.Context) (*entities.User, error)
	FindByCpf(cpf string, ctx context.Context) (*entities.User, error)
	Update(user *entities.User, ctx context.Context) error
	Delete(id string, ctx context.Context) error
}

type UserRepositoryDb struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepositoryDb {
	return &UserRepositoryDb{

		conn: conn,
	}
}

func (r *UserRepositoryDb) Insert(user *entities.User, ctx context.Context) error {
	if user.ID.String() == "" {
		user.ID = pkg_entities.NewID()
	}
	sql := `INSERT INTO users (id, name, last_name, email, cpf, password, created_at, update_log) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {

		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.Name,
			user.LastName,
			user.Email,
			user.CPF,
			user.Password,
			user.CreatedAt,
			user.UpdateLog,
		)

		if err != nil {
			return errors.ErrCtx(err, "r.tx.Exec")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) FindByEmail(email string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE email = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, email)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) FindById(id pkg_entities.ID, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE id = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, id)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) FindByCpf(cpf string, ctx context.Context) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE cpf = $1`
	var row pgx.Row
	user := &entities.User{}
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		row = tx.QueryRow(ctx, sql, cpf)
		err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.CPF, &user.Password, &user.CreatedAt, &user.UpdateLog)
		if err != nil {
			return errors.ErrCtx(err, "row.Scan")
		}
		return nil
	})
	if err != nil {
		return nil, errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return user, nil
}

func (r *UserRepositoryDb) Update(user *entities.User, ctx context.Context) error {
	sql := `UPDATE users SET name = $2, last_name = $3, email = $4, cpf = $5, password = $6, update_log = $7 WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql,
			user.ID,
			user.Name,
			user.LastName,
			user.Email,
			user.CPF,
			user.Password,		
			user.UpdateLog,
		)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec(")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}

func (r *UserRepositoryDb) Delete(id string, ctx context.Context) error {
	sql := `DELETE FROM users WHERE id = $1`
	err := crdbpgx.ExecuteTx(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, sql, id)
		if err != nil {
			return errors.ErrCtx(err, "tx.Exec(")
		}
		return nil
	})
	if err != nil {
		return errors.ErrCtx(err, "crdbpgx.ExecuteTx")
	}
	return nil
}
