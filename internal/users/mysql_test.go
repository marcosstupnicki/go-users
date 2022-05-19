package users

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMySQL_Create(t *testing.T) {
	user := User{
		Email:    "some@email.com",
		Password: "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G",
	}

	var tests = []struct {
		name           string
		user           User
		db             *gorm.DB
		expectedResult User
		expectedError  error
	}{
		{
			name: "Ok - Create",
			user: user,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedResult: User{
				ID:        0,
				Email:     "some@email.com",
				Password:  "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		{
			name: "Fail - Internal error",
			user: user,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users`").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("internal error"))
				mock.ExpectCommit()
				mock.ExpectRollback()

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := MySQL{
				DB: tt.db,
			}
			result, err := repo.Create(tt.user)

			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestMySQL_Get(t *testing.T) {
	var tests = []struct {
		name           string
		id             int
		db             *gorm.DB
		expectedResult User
		expectedError  error
	}{
		{
			name: "Ok - Get user",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
					AddRow(1, "some@email.com", "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G", 123456, 123456)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedResult: User{
				ID: 1,
				Email: "some@email.com",
				Password: "$2a$10$i8u5FgiJXRui/p.ZDXnDO.kVq3H6rbqrQp6rInFX.IeEO0zN/2F5G",
				CreatedAt: 123456,
				UpdatedAt: 123456,
			},
		},
		{
			name: "Fail - User not found",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				row := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"})

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(row)

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedError: ErrUserNotFound,
		},
		{
			name: "Fail - Internal error",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(errors.New("internal error"))

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedError: ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := MySQL{
				DB: tt.db,
			}
			result, err := repo.Get(tt.id)

			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestMySQL_Delete(t *testing.T) {
	var tests = []struct {
		name           string
		id             int
		db             *gorm.DB
		expectedError  error
	}{
		{
			name: "Ok - Delete user",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
		},
		{
			name: "Fail - User not found",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectCommit()

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedError: ErrUserNotFound,
		},
		{
			name: "Fail - Internal error",
			id:   1,
			db: func() *gorm.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)

				mock.MatchExpectationsInOrder(false)
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("DELETE")).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(errors.New("internal error"))
				mock.ExpectRollback()

				gormDB, err := gorm.Open(
					mysql.New(mysql.Config{
						Conn:                      db,
						SkipInitializeWithVersion: true}),
					&gorm.Config{})
				require.NoError(t, err)

				return gormDB
			}(),
			expectedError: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := MySQL{
				DB: tt.db,
			}
			err := repo.Delete(tt.id)

			require.Equal(t, tt.expectedError, err)
		})
	}
}