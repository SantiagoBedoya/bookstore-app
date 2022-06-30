package cassandra

import (
	accesstoken "github.com/SantiagoBedoya/bookstore_oauth-api/access_token"
	"github.com/SantiagoBedoya/bookstore_utils/logger"
	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"github.com/gocql/gocql"
)

const (
	GET_ACCESS_TOKEN_QUERY    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
	CREATE_ACCESS_TOKEN_QUERY = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	UPDATE_ACCESS_TOKEN_QUERY = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

type repository struct {
	session *gocql.Session
}

func newCassandraClient(host string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

func NewCassandraRepository(host string) (accesstoken.Repository, *gocql.Session) {
	session, err := newCassandraClient(host)
	if err != nil {
		panic(err)
	}
	return &repository{
		session: session,
	}, session
}

func (r *repository) GetAccessTokenById(id string) (*accesstoken.AccessToken, *resterrors.RestError) {
	var at accesstoken.AccessToken
	if err := r.session.Query(GET_ACCESS_TOKEN_QUERY, id).Scan(
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	); err != nil {
		logger.Error(err.Error())
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("access token not found")
		}
		return nil, resterrors.NewInternalServerError(accesstoken.GetAccessTokenErr.Error())
	}
	return &at, nil
}

func (r *repository) CreateAccessToken(at *accesstoken.AccessToken) (*accesstoken.AccessToken, *resterrors.RestError) {
	if err := r.session.Query(
		CREATE_ACCESS_TOKEN_QUERY,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewInternalServerError(accesstoken.CreateAccessTokenErr.Error())
	}
	return at, nil
}

func (r *repository) UpdateAccessTokenExpirationTime(at *accesstoken.AccessToken) *resterrors.RestError {
	if err := r.session.Query(
		UPDATE_ACCESS_TOKEN_QUERY,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		logger.Error(err.Error())
		return resterrors.NewInternalServerError(accesstoken.UpdateAccessTokenErr.Error())
	}
	return nil
}
