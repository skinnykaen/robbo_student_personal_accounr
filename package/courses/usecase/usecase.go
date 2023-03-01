package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"log"
)

type CourseUseCaseImpl struct {
	courseGateway courses.Gateway
}

type CourseUseCaseModule struct {
	fx.Out
	courses.UseCase
}

func SetupCourseUseCase(gateway courses.Gateway) CourseUseCaseModule {
	return CourseUseCaseModule{
		UseCase: &CourseUseCaseImpl{
			courseGateway: gateway,
		},
	}
}

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseCore) (id string, err error) {
	CourseId, err := p.courseGateway.CreateCourse(course)
	if err != nil {
		log.Println("Error create Course")
		return "", err
	}

	mediaCore := &models.CourseApiMediaCollectionCore{
		CourseID: CourseId,
	}
	MediaId, err := p.courseGateway.CreateCourseApiMediaCollection(mediaCore)
	if err != nil {
		log.Println("Error create CourseApiMediaCollection")
		return "", err
	}

	bannerImage := &models.AbsoluteMediaCore{
		Uri:                        course.Media.BannerImage.Uri,
		UriAbsolute:                course.Media.BannerImage.UriAbsolute,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.courseGateway.CreateAbsoluteMedia(bannerImage)
	if err != nil {
		log.Println("Error create AbsoluteMedia")
		return "", err
	}

	courseImage := &models.MediaCore{
		Uri:                        course.Media.CourseImage.Uri,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.courseGateway.CreateMedia(courseImage)
	if err != nil {
		log.Println("Error create Media")
		return "", err
	}

	courseVideo := &models.MediaCore{
		Uri:                        course.Media.CourseVideo.Uri,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.courseGateway.CreateMedia(courseVideo)
	if err != nil {
		log.Println("Error create Media")
		return "", err
	}

	image := &models.ImageCore{
		Raw:                        course.Media.Image.Raw,
		Small:                      course.Media.Image.Small,
		Large:                      course.Media.Image.Large,
		CourseApiMediaCollectionID: MediaId,
	}
	_, err = p.courseGateway.CreateImage(image)
	if err != nil {
		log.Println("Error create Image")
		return "", err
	}

	return CourseId, nil
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {
	err = p.courseGateway.UpdateAbsoluteMedia(&course.Media.BannerImage)
	if err != nil {
		log.Println("Error update AbsoluteMedia")
		return
	}

	err = p.courseGateway.UpdateMedia(&course.Media.CourseImage)
	if err != nil {
		log.Println("Error update Media")
		return
	}

	err = p.courseGateway.UpdateMedia(&course.Media.CourseVideo)
	if err != nil {
		log.Println("Error update Media")
		return
	}

	err = p.courseGateway.UpdateImage(&course.Media.Image)
	if err != nil {
		log.Println("Error update Image")
		return
	}

	err = p.courseGateway.UpdateCourseApiMediaCollection(&course.Media)
	if err != nil {
		log.Println("Error update CourseApiMediaCollection")
		return
	}

	err = p.courseGateway.UpdateCourse(course)
	if err != nil {
		log.Println("Error update Course")
		return
	}

	return nil
}

func (p *CourseUseCaseImpl) GetAccessCourseRelationsByStudentId(
	studentId string,
) (courseRelations []*models.CourseRelationCore, err error) {
	return p.courseGateway.GetAccessCourseRelationsByStudentId(studentId)
}

func (p *CourseUseCaseImpl) DeleteCourse(courseId string) (err error) {
	id, err := p.courseGateway.DeleteCourse(courseId)
	if err != nil {
		log.Println("Error delete Course")
		return
	}

	courseApiMediaCollectionId, err := p.courseGateway.DeleteCourseApiMediaCollection(id)
	if err != nil {
		log.Println("Error delete CourseApiMediaCollection")
		return
	}

	err = p.courseGateway.DeleteAbsoluteMedia(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete AbsoluteMedia")
		return
	}

	err = p.courseGateway.DeleteMedia(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete Media")
		return
	}

	err = p.courseGateway.DeleteImage(courseApiMediaCollectionId)
	if err != nil {
		log.Println("Error delete Image")
		return
	}

	return nil
}
