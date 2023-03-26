package delegate

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type CourseDelegateImpl struct {
	CoursesUseCase courses.UseCase
	EdxUseCase     edx.UseCase
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(coursesUsecase courses.UseCase, edxUsecase edx.UseCase) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{
			coursesUsecase,
			edxUsecase,
		},
	}
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	body, err := p.EdxUseCase.GetCourseContent(courseId)
	if err != nil {
		return "", courses.ErrBadRequest
	}
	err = json.Unmarshal(body, course)
	if err != nil {
		return "", courses.ErrInternalServerLevel
	}
	courseCore := course.ToCore()
	return p.CoursesUseCase.CreateCourse(courseCore)
}

func (p *CourseDelegateImpl) DeleteCourse(courseId string) (err error) {
	return p.CoursesUseCase.DeleteCourse(courseId)
}

func (p *CourseDelegateImpl) UpdateCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.CoursesUseCase.UpdateCourse(courseCore)
}

func (p *CourseDelegateImpl) GetCourseContent(courseId string) (courseHTTP *models.CourseHTTP, err error) {
	body, err := p.EdxUseCase.GetCourseContent(courseId)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &courseHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return courseHTTP, nil
}

func (p *CourseDelegateImpl) GetCoursesByUser(userId string, role models.Role, page string, pageSize string) (
	coursesListHTTP *models.CoursesListHTTP,
	err error,
) {
	var courseAccessRelations []*models.CourseRelationCore
	var errGetRelations error
	fmt.Println(role)
	switch role {
	case models.Student:
		courseAccessRelations, errGetRelations = p.CoursesUseCase.GetAccessCourseRelationsByStudentId(userId)
		break
	case models.SuperAdmin:
		body, err := p.EdxUseCase.GetCoursesByUser()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		err = json.Unmarshal(body, &coursesListHTTP)
		if err != nil {
			return nil, courses.ErrInternalServerLevel
		}
		return coursesListHTTP, nil
	}

	if errGetRelations != nil {
		return nil, errGetRelations
	}
	for _, courseAccessRelation := range courseAccessRelations {
		var courseHTTP *models.CourseHTTP
		body, err := p.EdxUseCase.GetCourseContent(courseAccessRelation.CourseId)
		if err != nil {
			return nil, courses.ErrBadRequest
		}
		err = json.Unmarshal(body, &courseHTTP)
		if err != nil {
			return nil, courses.ErrInternalServerLevel
		}
		coursesListHTTP = &models.CoursesListHTTP{
			Results:    []*models.CourseHTTP{},
			Pagination: &models.Pagination{},
			CountRows:  0,
		}
		coursesListHTTP.Results = append(coursesListHTTP.Results, courseHTTP)
	}
	coursesListHTTP.CountRows = len(courseAccessRelations)
	return coursesListHTTP, nil
}

func (p *CourseDelegateImpl) GetEnrollments(username string) (enrollmentsListHTTP *models.EnrollmentsListHTTP, err error) {
	body, err := p.EdxUseCase.GetEnrollments(username)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &enrollmentsListHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return enrollmentsListHTTP, nil
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber string) (coursesListHTTP *models.CoursesListHTTP, err error) {
	pN, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	body, err := p.EdxUseCase.GetAllPublicCourses(pN)
	if err != nil {
		return nil, courses.ErrBadRequest
	}
	err = json.Unmarshal(body, &coursesListHTTP)
	if err != nil {
		return nil, courses.ErrInternalServerLevel
	}
	return coursesListHTTP, nil
}

func (p *CourseDelegateImpl) PostEnrollment(postEnrollmentHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxUseCase.PostEnrollment(postEnrollmentHTTP.Message)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxUseCase.PostEnrollment(postUnenrollHTTP.Message)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) Login(email, password string) (err error) {
	_, err = p.EdxUseCase.Login(email, password)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}

func (p *CourseDelegateImpl) Registration(userForm *edx.RegistrationForm) (err error) {
	_, err = p.EdxUseCase.PostRegistration(*userForm)
	if err != nil {
		log.Println(err)
		return courses.ErrBadRequest
	}
	return
}
