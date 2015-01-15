package main

/*
var (
	GOOGLE_OAUTH2_CLIENT_ID     = gqConfig.googleOauth2Id
	GOOGLE_OAUTH2_CLIENT_SECRET = gqConfig.googleOauth2Secret
)

type OpenIdResponse struct {
	Kind          string `json:"kind"`
	Gender        string `json:"gender"`
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Locale        string `json:"locale"`
	Hd            string `json:"hd"`
}

func (o *OpenIdResponse) IsGoogle() bool {
	return o.Kind == "plus#personOpenIdConnect"
}

func GetUserInfo(token oauth2.Tokens) *OpenIdResponse {
	oidc := new(OpenIdResponse)
	resp, err := http.Get(fmt.Sprint("https://www.googleapis.com/plus/v1/people/me/openIdConnect?access_token=", token.Access()))
	check(err)

	if err != nil {
		log.Panic("Error reading JSON response: ", err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	check(err)
	err = json.Unmarshal(contents, oidc)
	if err != nil {
		log.Panic("Error unmarshaling JSON response: ", err)
	}
	return oidc
}
*/
