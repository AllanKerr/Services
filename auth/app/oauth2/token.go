package oauth2

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/ory/fosite"
	"context"
)

func (ctrl *HTTPController) Token(w http.ResponseWriter, req *http.Request) {

	logger := logrus.WithFields(logrus.Fields{"endpoint": req.URL})
	logger.Debug("Handling request.")

	session := new(fosite.DefaultSession)
	ctx := context.Background()

	accessRequest, err := ctrl.auth.NewAccessRequest(ctx, req, session)
	if err != nil {
		logger.Warning(err)
		ctrl.auth.WriteAccessError(w, accessRequest, err)
		return
	} else if accessRequest.GetGrantTypes().Exact(PASSWORD_GRANT) {
		session.Username = accessRequest.GetRequestForm().Get("username")
	}

	response, err := ctrl.auth.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		logger.Warning(err)
		ctrl.auth.WriteAccessError(w, accessRequest, err)
		return
	}
	ctrl.auth.WriteAccessResponse(w, accessRequest, response)
}
