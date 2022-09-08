package mysql

import (
	"context"
	"database/sql"
	"deall/cmd/entity"
	"time"

	"github.com/sirupsen/logrus"
)

type MySQL struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *MySQL {
	return &MySQL{db}
}

func (m *MySQL) ListUser(ctx context.Context, page, limit int64) ([]*entity.User, error) {
	var users []*entity.User
	var offset = (page - 1) * limit

	query := `SELECT
		id,
		username,
		email,
		password,
		first_name,
		last_name,
		phone_number,
		role_id,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM tm_user
	WHERE deleted_at IS NULL
		AND deleted_by IS NULL
	ORDER BY created_at DESC
	LIMIT ?, ?;`

	rows, err := m.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(entity.User)
		err := rows.Scan(
			&row.ID,
			&row.Username,
			&row.Email,
			&row.Password,
			&row.FirstName,
			&row.LastName,
			&row.PhoneNumber,
			&row.RoleID,
			&row.CreatedAt,
			&row.CreatedBy,
			&row.UpdatedAt,
			&row.UpdatedBy,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		users = append(users, row)
	}
	return users, nil
}

func (m *MySQL) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User

	query := `
	SELECT 
		id,
		username,
		email,
		password,
		first_name,
		last_name,
		phone_number,
		role_id,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM tm_user
	WHERE id = ?
		AND deleted_at IS NULL
		AND deleted_by IS NULL;`

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.RoleID,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &user, nil
}

func (m *MySQL) GetUserRoleID(ctx context.Context, roleID string) ([]*entity.User, error) {
	var users []*entity.User

	query := `SELECT
		id,
		username,
		email,
		password,
		first_name,
		last_name,
		phone_number,
		role_id,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM tm_user
	WHERE role_id = ?
		AND deleted_at IS NULL
		AND deleted_by IS NULL
	ORDER BY created_at DESC;`

	rows, err := m.db.QueryContext(ctx, query, roleID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(entity.User)
		err := rows.Scan(
			&row.ID,
			&row.Username,
			&row.Email,
			&row.Password,
			&row.FirstName,
			&row.LastName,
			&row.PhoneNumber,
			&row.RoleID,
			&row.CreatedAt,
			&row.CreatedBy,
			&row.UpdatedAt,
			&row.UpdatedBy,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		users = append(users, row)
	}
	return users, nil
}

func (m *MySQL) GetUserByEmailOrPhone(ctx context.Context, data string) (*entity.User, error) {
	var user entity.User

	query := `
	SELECT 
		id,
		username,
		email,
		password,
		first_name,
		last_name,
		phone_number,
		role_id,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM tm_user
	WHERE email = ? 
		OR phone_number = ?
		AND deleted_at IS NULL
		AND deleted_by IS NULL;`

	err := m.db.QueryRowContext(ctx, query, data, data).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.RoleID,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &user, nil
}

func (m *MySQL) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.CreatedAt = time.Now()
	query := `INSERT INTO tm_user
		(
			id,
			username,
			email,
			password,
			first_name,
			last_name,
			phone_number,
			role_id,
			created_at,
			created_by
		)
		VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.RoleID,
		user.CreatedAt,
		user.CreatedBy,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return user, nil
}

func (m *MySQL) UpdateUser(ctx context.Context, user *entity.User, id string) (*entity.User, error) {

	var now = time.Now()
	user.UpdatedAt = &now
	query := `UPDATE tm_user SET
		username = ?,
		email = ?,
		password = ?,
		first_name = ?,
		last_name = ?,
		phone_number = ?,
		role_id = ?,
		updated_at = ?,
		updated_by = ?
	WHERE id = ?
	`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.RoleID,
		user.UpdatedAt,
		user.UpdatedBy,
		id,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return user, nil
}

func (m *MySQL) DeleteUser(ctx context.Context, id string, userID string) error {
	query := `UPDATE tm_user SET
		deleted_at = NOW(),
		deleted_by = ?
	WHERE id  = ?
	`

	_, err := m.db.ExecContext(ctx, query, userID, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
