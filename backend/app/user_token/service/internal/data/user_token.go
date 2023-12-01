package data

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/devexps/go-micro/middleware/authn/engine/jwt/v2"

	"github.com/devexps/go-micro/v2/log"

	userV1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
	v1 "github.com/devexps/go-microservices-demo/gen/api/go/user_token/service/v1"
)

const (
	userAccessTokenKeyPrefix  = "a_uat_"
	userRefreshTokenKeyPrefix = "a_urt_"
)

// UserTokenRepo .
type UserTokenRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserTokenRepo .
func NewUserTokenRepo(data *Data, logger log.Logger) *UserTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user_token/repo/user_token-service"))
	return &UserTokenRepo{
		data: data,
		log:  l,
	}
}

// GenerateToken .
func (r *UserTokenRepo) GenerateToken(ctx context.Context, user *userV1.User) (*v1.TokenResponse, error) {
	accessToken := r.createAccessJwtToken(user.GetId())
	if accessToken == "" {
		return nil, v1.ErrorCreateAccessTokenFailed("create access token failed")
	}
	if err := r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		r.log.Errorf("GenerateToken store to redis failed, err=%v", err)
		return nil, v1.ErrorStoreRedisFailed("store to redis failed")
	}
	refreshToken := r.createRefreshToken()
	if refreshToken == "" {
		return nil, v1.ErrorCreateRefreshTokenFailed("create refresh token failed")
	}
	if err := r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		r.log.Errorf("GenerateToken store to redis failed, err=%v", err)
		return nil, v1.ErrorStoreRedisFailed("store to redis failed")
	}
	return &v1.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *UserTokenRepo) RemoveToken(ctx context.Context, user *userV1.User) (*emptypb.Empty, error) {
	if err := r.deleteAccessTokenFromRedis(ctx, user.GetId()); err != nil {
		return nil, v1.ErrorRemoveAccessTokenFailed("remove access token failed")
	}
	if err := r.deleteRefreshTokenFromRedis(ctx, user.GetId()); err != nil {
		return nil, v1.ErrorRemoveRefreshTokenFailed("remove refresh token failed")
	}
	return &emptypb.Empty{}, nil
}

func (r *UserTokenRepo) GetRefreshToken(ctx context.Context, user *userV1.User) (*v1.TokenResponse, error) {
	return &v1.TokenResponse{
		RefreshToken: r.getRefreshTokenFromRedis(ctx, user.GetId()),
	}, nil
}

func (r *UserTokenRepo) GenerateAccessToken(ctx context.Context, user *userV1.User) (*v1.TokenResponse, error) {
	accessToken := r.createAccessJwtToken(user.GetId())
	if accessToken == "" {
		return nil, v1.ErrorCreateAccessTokenFailed("create access token failed")
	}
	if err := r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		r.log.Errorf("GenerateAccessToken store to redis failed, err=%v", err)
		return nil, v1.ErrorStoreRedisFailed("store to redis failed")
	}
	return &v1.TokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (r *UserTokenRepo) createAccessJwtToken(userId uint32) string {
	principal := &jwt.Claims{
		Subject: strconv.FormatUint(uint64(userId), 10),
	}
	signedToken, err := r.data.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}
	return signedToken
}

func (r *UserTokenRepo) createRefreshToken() string {
	strUUID, _ := uuid.NewV4()
	return strUUID.String()
}

func (r *UserTokenRepo) setAccessTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", userAccessTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserTokenRepo) setRefreshTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserTokenRepo) deleteAccessTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userAccessTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}

func (r *UserTokenRepo) deleteRefreshTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}

func (r *UserTokenRepo) getRefreshTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("getRefreshTokenFromRedis get redis user refresh token failed, err=%v", err)
		}
		return ""
	}
	return result
}
