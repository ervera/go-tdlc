package handler

import (
	"github.com/ervera/tdlc-gin/internal/localGoogle"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "505960415454-ov0v0c2rh6hchkedgkrq6rqgjm3n3nad.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

type GoogleHandler struct {
	service localGoogle.Service
}

func NewGoogleHandler(p localGoogle.Service) *GoogleHandler {
	return &GoogleHandler{
		service: p,
	}
}

// func (c *GoogleHandler) Login() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token := ctx.Param("token")
// 		result, err := c.service.Login(ctx, token)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, result)
// 		return
// 	}
// }

// func HandleGoogleLogin() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		url := googleOauthConfig.AuthCodeURL(oauthStateString)
// 		web.Response(ctx, 200, url)
// 	}
// }

// func HandleGoogleCallback() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token := ctx.Param("token")
// 		result, err := GetPostUserInfo(token)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 		}
// 		web.Response(ctx, 200, result)
// 		//ctx.Redirect(200, "http://localhost:8080/"+string(content))
// 	}
//}

// func GetPostUserInfo(token string) (domain.GoogleUser, error) {
// 	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
// 	var googleResp domain.GoogleUser
// 	if err != nil {
// 		return googleResp, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}

// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return googleResp, fmt.Errorf("failed reading response body: %s", err.Error())
// 	}
// 	err = json.Unmarshal(contents, &googleResp)
// 	if err != nil {
// 		return googleResp, fmt.Errorf("failed unmarshal contents: %s", err.Error())
// 	}
// 	if googleResp.ID == "" {
// 		return domain.GoogleUser{}, errors.New("no google id")
// 	}
// 	return googleResp, nil
// }
