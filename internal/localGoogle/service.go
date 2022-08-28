package localGoogle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/ervera/tdlc-gin/pkg/random"
)

type Service interface {
	Login(ctx context.Context, token string) (domain.User, error)
}

type service struct {
	userService    user.Service
	userRepository user.Repository
}

func NewService(r user.Repository, s user.Service) Service {
	return &service{
		userService:    s,
		userRepository: r,
	}
}

func (s *service) Login(ctx context.Context, token string) (domain.User, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var googleResp domain.GoogleUser
	err = json.Unmarshal(contents, &googleResp)
	fmt.Println(googleResp)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed unmarshal contents: %s", err.Error())
	}

	if googleResp.ID == "" {
		return domain.User{}, errors.New("no google id")
	}

	user, exist := s.googleLogin(ctx, googleResp.Email)
	if exist {
		return user, nil
	}

	_, err = s.userService.CreateUser(ctx, googleUserToUser(googleResp))
	if err != nil {
		return domain.User{}, err
	}
	user, exist = s.googleLogin(ctx, googleResp.Email)
	if exist {
		return user, nil
	}

	token, ok := jwt.GenerateJWT(user)
	if ok != nil {
		return domain.User{}, nil
	}

	user.Token = token
	return user, nil
}

func (s *service) googleLogin(ctx context.Context, email string) (domain.User, bool) {
	user, err := s.userRepository.ExistAndGetByMail(ctx, email)
	if err != nil {
		return user, false
	}
	user.Password = ""

	token, ok := jwt.GenerateJWT(user)
	if ok != nil {
		return domain.User{}, false
	}

	user.Token = token
	return user, true
}

func googleUserToUser(g domain.GoogleUser) domain.User {
	return domain.User{
		FirstName: g.GivenName,
		LastName:  g.FamilyName,
		Avatar:    g.Picture,
		Email:     g.Email,
		Password:  random.GenerateStringByN(7),
	}
}
