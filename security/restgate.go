package security

/*
|--------------------------------------------------------------------------
| WARNING
|--------------------------------------------------------------------------
| Always use this middleware library with a HTTPS Connection.
| The Key and Password will be exposed and highly unsecure otherwise!
| The database server should also use HTTPS Connection and be hidden away
|
*/

/*
Thanks to Ido Ben-Natan ("IdoBn") for postgres fix.
*/

import (
	"github.com/gorilla/context"
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/models"
	"net/http"
)

type RESTGate struct {
	HeaderKeyLabel    string
	HeaderSecretLabel string
	UserIDKey         int
}

func New(headerKeyLabel string, headerSecretLabel string, userIdKey int) *RESTGate {

	if headerKeyLabel == "" {
		headerKeyLabel = "X-Auth-Key"
	}

	if headerSecretLabel == "" {
		headerSecretLabel = "X-Auth-Secret"
	}

	t := &RESTGate{HeaderKeyLabel: headerKeyLabel, HeaderSecretLabel: headerSecretLabel, UserIDKey: userIdKey}

	return t
}

func (self *RESTGate) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	//Check key in Header
	key := req.Header.Get(self.HeaderKeyLabel)
	secret := req.Header.Get(self.HeaderSecretLabel)

	if key == "" {
		//Authentication Information not included in request

		err := errors.Error{Description: "Error: UnAuthorized", Code: http.StatusUnauthorized}
		models.NewErrorOutput(err).Render(rw)
		return
	}

	if userid, err0 := models.Authenticate(key, secret); err0 == nil {
		//found the session
		context.Set(req, self.UserIDKey, userid)
		next(rw, req)
	} else {
		err := errors.Error{Description: "Error: UnAuthorized", Code: http.StatusUnauthorized}
		models.NewErrorOutput(err).Render(rw)
		return
	}

}
