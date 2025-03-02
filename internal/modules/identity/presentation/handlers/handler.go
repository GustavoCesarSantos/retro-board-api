package identity

type Handlers struct {
	RefreshAuthToken *RefreshAuthToken
	SigninUser *SigninUser
	SigninWithGoogle *SigninWithGoogle
	SigninWithGoogleCallback *SigninWithGoogleCallback
	SignoutUser *SignoutUser
}

func NewHandlers(
	refreshAuthToken *RefreshAuthToken,
	signinUser *SigninUser,
	signinWithGoogle *SigninWithGoogle,
	signinWithGoogleCallback *SigninWithGoogleCallback,
	signoutUser *SignoutUser,
) *Handlers {
	return &Handlers{
		RefreshAuthToken: refreshAuthToken,
		SigninUser: signinUser,
		SigninWithGoogle: signinWithGoogle,
		SigninWithGoogleCallback: signinWithGoogleCallback,
		SignoutUser: signoutUser,
	}
}