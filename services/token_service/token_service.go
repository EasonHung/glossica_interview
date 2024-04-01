package token_service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_EXPIRE_TIME = 60
const JWT_SECRET = "jwt-secret"

type JwtClaim struct {
	UserId          string    `json:"userId"`
	ExpireTime      time.Time `json:"expireTime"`
	SubmisstionTime time.Time `json:"submisstionTime"`
	RefreshToken    bool      `json:"refreshToken"`
}

func (JwtClaim) Valid() error {
	return nil
}

func CreateAccessToken(userId string) (error, string) {
	expireTime := shiftMinutes(ACCESS_TOKEN_EXPIRE_TIME).UTC()
	claim := JwtClaim{
		UserId:          userId,
		ExpireTime:      expireTime,
	}
	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, err := accessTokenClaims.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return errors.New("generate access token failed!"), ""
	}

	return nil, accessToken
}

func VerifyAccessToken(accessToken string) (error, *JwtClaim) {
	claim, err := getClaimFromJwt(accessToken)
	if err != nil {
		return errors.New("wrong token"), nil
	}

	// todo:
	// Use redis compare token submission time and user last login time.
	// If token submission time is earlier than user last login time, return error.

	if claim.ExpireTime.Before(time.Now().UTC()) {
		return errors.New("token expired"), nil
	}

	return nil, claim
}

func shiftMinutes(minutes int) time.Time {
	return time.Now().Local().Add(time.Minute * time.Duration(minutes))
}

func getClaimFromJwt(token string) (*JwtClaim, error) {
	var claim *JwtClaim
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		err = errors.New("error get token claim")
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*JwtClaim); ok {
		claim = claims
	}

	return claim, nil
}