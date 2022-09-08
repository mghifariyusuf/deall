package mysql

import (
	"context"
	"deall/cmd/entity"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
)

var createdBy = "o"
var userPayload = []*entity.User{
	&entity.User{
		ID:          "ini id",
		Username:    "user1",
		Email:       "coba@mail.com",
		FirstName:   "first_name",
		LastName:    "last_name",
		Password:    "ok",
		RoleID:      "ok",
		PhoneNumber: "0813456789",
		CreatedAt:   time.Now(),
		CreatedBy:   "nil",
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	},
	&entity.User{
		ID:          "ini idi",
		Username:    "user2",
		Email:       "id2@email.com",
		FirstName:   "first2_name",
		LastName:    "last2_name",
		Password:    "ok",
		RoleID:      "ok",
		PhoneNumber: "08134567890",
		CreatedAt:   time.Now(),
		CreatedBy:   "nil",
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	},
}

func TestListUser(t *testing.T) {
	var limit int64 = 10
	var page int64 = 1
	var offset = (page - 1) * limit
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	defer db.Close()

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

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id",
			"username",
			"email",
			"first_name",
			"last_name",
			"password",
			"phone_number",
			"role_id",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		})

		rows.AddRow(
			userPayload[0].ID,
			userPayload[0].Username,
			userPayload[0].Email,
			userPayload[0].FirstName,
			userPayload[0].LastName,
			userPayload[0].Password,
			userPayload[0].PhoneNumber,
			userPayload[0].RoleID,
			userPayload[0].CreatedAt,
			userPayload[0].CreatedBy,
			userPayload[0].UpdatedAt,
			userPayload[0].UpdatedBy,
		).AddRow(
			userPayload[1].ID,
			userPayload[1].Username,
			userPayload[1].Email,
			userPayload[1].FirstName,
			userPayload[1].LastName,
			userPayload[1].Password,
			userPayload[1].PhoneNumber,
			userPayload[1].RoleID,
			userPayload[1].CreatedAt,
			userPayload[1].CreatedBy,
			userPayload[1].UpdatedAt,
			userPayload[1].UpdatedBy,
		)

		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(offset, limit).
			WillReturnRows(rows)

		repo := NewUser(db)
		userResponse, err := repo.ListUser(context.TODO(), page, limit)
		require.NoError(t, err)
		require.NotEmpty(t, userResponse)
		require.Len(t, userResponse, len(userPayload))
		require.Equal(t, userResponse[0].ID, userPayload[0].ID)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(offset, limit).
			WillReturnError(fmt.Errorf("Error"))

		repo := NewUser(db)
		userResponse, err := repo.ListUser(context.TODO(), page, limit)
		require.Nil(t, userResponse)
		require.Error(t, err)
	})
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not exptected when opening stub database connection", err)
	}

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

	t.Run("success", func(t *testing.T) {
		row := sqlmock.NewRows([]string{
			"id",
			"username",
			"email",
			"first_name",
			"last_name",
			"password",
			"phone_number",
			"role_id",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by"}).
			AddRow(userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].Password,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy)

		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(userPayload[0].ID).
			WillReturnRows(row)

		repo := NewUser(db)
		userResponse, err := repo.GetUserByID(context.TODO(), userPayload[0].ID)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		require.Equal(t, userResponse.ID, userPayload[0].ID)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(userPayload[0].ID).
			WillReturnError(errors.New("error"))

		repo := NewUser(db)
		userResponse, err := repo.GetUserByID(context.TODO(), userPayload[0].ID)
		require.Error(t, err)
		require.Nil(t, userResponse)
	})
}

func TestGetUserByRoleID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	defer db.Close()

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

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id",
			"username",
			"email",
			"first_name",
			"last_name",
			"password",
			"phone_number",
			"role_id",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		})

		rows.AddRow(
			userPayload[0].ID,
			userPayload[0].Username,
			userPayload[0].Email,
			userPayload[0].FirstName,
			userPayload[0].LastName,
			userPayload[0].Password,
			userPayload[0].PhoneNumber,
			userPayload[0].RoleID,
			userPayload[0].CreatedAt,
			userPayload[0].CreatedBy,
			userPayload[0].UpdatedAt,
			userPayload[0].UpdatedBy,
		).AddRow(
			userPayload[1].ID,
			userPayload[1].Username,
			userPayload[1].Email,
			userPayload[1].FirstName,
			userPayload[1].LastName,
			userPayload[1].Password,
			userPayload[1].PhoneNumber,
			userPayload[1].RoleID,
			userPayload[1].CreatedAt,
			userPayload[1].CreatedBy,
			userPayload[1].UpdatedAt,
			userPayload[1].UpdatedBy,
		)

		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(userPayload[0].RoleID).
			WillReturnRows(rows)

		repo := NewUser(db)
		userResponse, err := repo.GetUserRoleID(context.TODO(), userPayload[0].RoleID)
		require.NoError(t, err)
		require.NotEmpty(t, userResponse)
		require.Len(t, userResponse, len(userPayload))
		require.Equal(t, userResponse[0].ID, userPayload[0].ID)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(userPayload[0].ID).
			WillReturnError(fmt.Errorf("error"))

		repo := NewUser(db)
		userResponse, err := repo.GetUserRoleID(context.TODO(), userPayload[0].RoleID)
		require.Error(t, err)
		require.Nil(t, userResponse)
	})
}

func TestGetUserByEmailOrPhone(t *testing.T) {
	var email = "coba@mail.com"
	var phoneNumber = "0813456789"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not exptected when opening stub database connection", err)
	}
	defer db.Close()

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

	t.Run("success-by-phone-number", func(t *testing.T) {
		row := sqlmock.NewRows([]string{
			"id",
			"username",
			"email",
			"first_name",
			"last_name",
			"password",
			"phone_number",
			"role_id",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by"}).
			AddRow(userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].Password,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy)

		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(phoneNumber, phoneNumber).
			WillReturnRows(row)
		repo := NewUser(db)
		userResponse, err := repo.GetUserByEmailOrPhone(context.TODO(), phoneNumber)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		require.Equal(t, userResponse.PhoneNumber, phoneNumber)
	})

	t.Run("success-by-email", func(t *testing.T) {
		row := sqlmock.NewRows([]string{
			"id",
			"username",
			"email",
			"first_name",
			"last_name",
			"password",
			"phone_number",
			"role_id",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by"}).
			AddRow(userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].Password,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy)

		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(email, email).
			WillReturnRows(row)
		repo := NewUser(db)
		userResponse, err := repo.GetUserByEmailOrPhone(context.TODO(), email)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
		require.Equal(t, userResponse.Email, email)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(email, email).
			WillReturnError(fmt.Errorf("Error"))
		repo := NewUser(db)
		userResponse, err := repo.GetUserByEmailOrPhone(context.TODO(), email)
		require.Error(t, err)
		require.Nil(t, userResponse)
	})
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not exptected when opening stub database connection", err)
	}
	defer db.Close()

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

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).
			ExpectExec().
			WithArgs(
				userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewUser(db)

		userResponse, err := repo.InsertUser(context.TODO(), userPayload[0])
		require.NoError(t, err)
		require.Equal(t, userPayload[0].ID, userResponse.ID)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).
			ExpectExec().
			WithArgs(
				userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
			).
			WillReturnError(fmt.Errorf("error"))

		repo := NewUser(db)

		userResponse, err := repo.InsertUser(context.TODO(), userPayload[0])
		require.Error(t, err)
		require.Nil(t, userResponse)
	})

	t.Run("error-statement", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(
				userPayload[0].ID,
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].CreatedAt,
				userPayload[0].CreatedBy,
			).
			WillReturnError(fmt.Errorf("error"))

		repo := NewUser(db)

		userResponse, err := repo.InsertUser(context.TODO(), userPayload[0])
		require.Error(t, err)
		require.Nil(t, userResponse)
	})
}

func TestUpdateUser(t *testing.T) {
	id := userPayload[0].ID

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not exptected when opening stub database connection", err)
	}
	defer db.Close()

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

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).
			ExpectExec().
			WithArgs(
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy,
				id,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewUser(db)
		userResponse, err := repo.UpdateUser(context.TODO(), userPayload[0], id)
		require.NoError(t, err)
		require.NotNil(t, userResponse)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).
			ExpectExec().
			WithArgs(
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy,
				id,
			).
			WillReturnError(fmt.Errorf("error"))

		repo := NewUser(db)
		userResponse, err := repo.UpdateUser(context.TODO(), userPayload[0], id)
		require.Error(t, err)
		require.Nil(t, userResponse)
	})

	t.Run("error-statement", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(
				userPayload[0].Username,
				userPayload[0].Email,
				userPayload[0].Password,
				userPayload[0].FirstName,
				userPayload[0].LastName,
				userPayload[0].PhoneNumber,
				userPayload[0].RoleID,
				userPayload[0].UpdatedAt,
				userPayload[0].UpdatedBy,
				id,
			).
			WillReturnError(fmt.Errorf("error"))

		repo := NewUser(db)
		userResponse, err := repo.UpdateUser(context.TODO(), userPayload[0], id)
		require.Error(t, err)
		require.Nil(t, userResponse)
	})
}

func TestDeleteUser(t *testing.T) {
	id := "ini id"
	userID := "nil"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not exptected when opening stub database connection", err)
	}

	query := `UPDATE tm_user SET
		deleted_at = NOW(),
		deleted_by = ?,
	WHERE id  = ?
	`

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(userID, id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewUser(db)
		err = repo.DeleteUser(context.TODO(), id, userID)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(userID, id).
			WillReturnError(errors.New("error"))

		repo := NewUser(db)
		err = repo.DeleteUser(context.TODO(), id, userID)
		require.Error(t, err)
	})
}
