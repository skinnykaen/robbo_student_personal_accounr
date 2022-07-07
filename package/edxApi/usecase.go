package edxApi

//go:generate mockgen -source=api_service.go -destination=mocks/mock.go

type NewToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type RegistrationForm struct {
	Email            string
	Username         string
	Name             string
	Password         string
	Terms_of_service string
}

type EdxApiUseCase interface {
	GetCoursesByUser() (respBody string, err error)
	GetAllPublicCourses(pageNumber int) (respBody string, err error)
	GetEnrollments(username string) (respBody string, err error)
	GetUser() (respBody string, err error)
	GetCourseContent(courseId string) (respBody string, err error)
	PostEnrollment(message map[string]interface{}) (respBody string, err error)
	PostRegistration(postMessage RegistrationForm) (respBody string, err error)
	Login(email, password string) (respBody string, err error)
}