package mysql

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/SantiagoBedoya/bookstore_users-api/users"
	"github.com/SantiagoBedoya/bookstore_utils/logger"
	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CREATE_USER_QUERY       = "INSERT INTO users(first_name, last_name, email, status, password, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	GET_USER_BY_ID_QUERY    = "SELECT * FROM users WHERE id = ?"
	GET_USER_BY_EMAIL_QUERY = "SELECT * FROM users WHERE email = ?"
	GET_USERS_BY_STATUS     = "SELECT * FROM users WHERE status = ?"
	UPDATE_USER_BY_ID       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	DELETE_USER_BY_ID       = "DELETE FROM users WHERE id = ?"
	GET_USERS               = "SELECT * FROM users"
)

type mysqlRepository struct {
	db *sql.DB
}

func newMysqlClient(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func NewMysqlRepository(connection string) users.Repository {
	repo := &mysqlRepository{}
	db, err := newMysqlClient(connection)
	if err != nil {
		log.Fatal(err)
	}
	repo.db = db
	return repo
}

func (r *mysqlRepository) CreateUser(user *users.User) (*users.User, *resterrors.RestError) {
	stmt, err := r.db.Prepare(CREATE_USER_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.SaveUserErr.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.CreatedAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			logger.Error(err.Error())
			return nil, resterrors.NewBadRequestError(users.UserAlreadyExistErr.Error())
		}
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.SaveUserErr.Error())
	}
	userId, err := result.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.SaveUserErr.Error())
	}
	user.Id = userId
	return user, nil
}

func (r *mysqlRepository) GetUser(userId string) (*users.User, *resterrors.RestError) {
	var user users.User
	stmt, err := r.db.Prepare(GET_USER_BY_ID_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUserErr.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(userId)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, resterrors.NewNotFoundError(users.UserNotFoundErr.Error())
		}
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUserErr.Error())
	}
	return &user, nil
}

func (r *mysqlRepository) SearchUserByEmail(email string) (*users.User, *resterrors.RestError) {
	var user users.User
	stmt, err := r.db.Prepare(GET_USER_BY_EMAIL_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUserErr.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(email)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, resterrors.NewNotFoundError(users.UserNotFoundErr.Error())
		}
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUserErr.Error())
	}
	return &user, nil
}

func (r *mysqlRepository) UpdateUser(userId string, user *users.User) (*users.User, *resterrors.RestError) {
	stmt, err := r.db.Prepare(UPDATE_USER_BY_ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.UpdateUserErr.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, userId)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			logger.Error(err.Error())
			return nil, resterrors.NewBadRequestError(users.UserAlreadyExistErr.Error())
		}
		return nil, resterrors.NewInternalServerError(users.UpdateUserErr.Error())
	}
	uid, _ := strconv.Atoi(userId)
	user.Id = int64(uid)
	return user, nil
}
func (r *mysqlRepository) DeleteUser(userId string) (*users.User, *resterrors.RestError) {
	stmt, err := r.db.Prepare(DELETE_USER_BY_ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.DeleteUserErr.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.DeleteUserErr.Error())
	}
	uid, _ := strconv.Atoi(userId)
	return &users.User{Id: int64(uid)}, nil
}

func (r *mysqlRepository) GetUsersByStatus(status string) ([]users.User, *resterrors.RestError) {
	stmt, err := r.db.Prepare(GET_USERS_BY_STATUS)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
	}
	results := make([]users.User, 0)
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.CreatedAt); err != nil {
			logger.Error(err.Error())
			return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
		}
		results = append(results, user)
	}
	return results, nil
}

func (r *mysqlRepository) GetUsers() ([]users.User, *resterrors.RestError) {
	stmt, err := r.db.Prepare(GET_USERS)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
	}
	results := make([]users.User, 0)
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.CreatedAt); err != nil {
			logger.Error(err.Error())
			return nil, resterrors.NewInternalServerError(users.GetUsersErr.Error())
		}
		results = append(results, user)
	}
	return results, nil
}
