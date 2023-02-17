// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type CohortResult interface {
	IsCohortResult()
}

type CourseRelationResult interface {
	IsCourseRelationResult()
}

type CourseRelationsResult interface {
	IsCourseRelationsResult()
}

type CourseResult interface {
	IsCourseResult()
}

type CoursesResult interface {
	IsCoursesResult()
}

type EnrollmentsResult interface {
	IsEnrollmentsResult()
}

type GetUserResult interface {
	IsGetUserResult()
}

type PairsStudentParentsResult interface {
	IsPairsStudentParentsResult()
}

type ParentResult interface {
	IsParentResult()
}

type ParentsResult interface {
	IsParentsResult()
}

type ProjectPageResult interface {
	IsProjectPageResult()
}

type RobboGroupResult interface {
	IsRobboGroupResult()
}

type RobboGroupsResult interface {
	IsRobboGroupsResult()
}

type RobboUnitResult interface {
	IsRobboUnitResult()
}

type RobboUnitsResult interface {
	IsRobboUnitsResult()
}

type SignInResult interface {
	IsSignInResult()
}

type StudentResult interface {
	IsStudentResult()
}

type StudentsResult interface {
	IsStudentsResult()
}

type SuperAdminResult interface {
	IsSuperAdminResult()
}

type TeacherResult interface {
	IsTeacherResult()
}

type TeachersResult interface {
	IsTeachersResult()
}

type UnitAdminResult interface {
	IsUnitAdminResult()
}

type UnitAdminsResult interface {
	IsUnitAdminsResult()
}

type AbsoluteMediaHTTP struct {
	ID          string `json:"id"`
	URI         string `json:"uri"`
	URIAbsolute string `json:"uri_absolute"`
}

type CohortHTTP struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	UserCount       int    `json:"user_count"`
	AssignmentType  string `json:"assignment_type"`
	UserPartitionID int    `json:"user_partition_id"`
	GroupID         int    `json:"group_id"`
}

func (CohortHTTP) IsCohortResult() {}

type CourseAPIMediaCollectionHTTP struct {
	ID          string             `json:"id"`
	BannerImage *AbsoluteMediaHTTP `json:"banner_image"`
	CourseImage *MediaHTTP         `json:"course_image"`
	CourseVideo *MediaHTTP         `json:"course_video"`
	Image       *ImageHTTP         `json:"image"`
}

type CourseHTTP struct {
	ID               string                        `json:"id"`
	BlocksURL        string                        `json:"blocks_url"`
	Effort           string                        `json:"effort"`
	EnrollmentStart  string                        `json:"enrollment_start"`
	EnrollmentEnd    string                        `json:"enrollment_end"`
	End              string                        `json:"end"`
	Name             string                        `json:"name"`
	Number           string                        `json:"number"`
	Org              string                        `json:"org"`
	ShortDescription string                        `json:"short_description"`
	Start            string                        `json:"start"`
	StartDisplay     string                        `json:"start_display"`
	StartType        string                        `json:"start_type"`
	Pacing           string                        `json:"pacing"`
	MobileAvailable  bool                          `json:"mobile_available"`
	Hidden           bool                          `json:"hidden"`
	InvitationOnly   bool                          `json:"invitation_only"`
	Overview         *string                       `json:"overview"`
	CourseID         string                        `json:"course_id"`
	Media            *CourseAPIMediaCollectionHTTP `json:"media"`
}

func (CourseHTTP) IsCourseResult() {}

type CourseRelationHTTP struct {
	ID           string `json:"id"`
	LastModified string `json:"lastModified"`
	Parameter    string `json:"parameter"`
	CourseID     string `json:"courseId"`
	ObjectID     string `json:"objectId"`
}

func (CourseRelationHTTP) IsCourseRelationResult() {}

type CourseRelationHTTPList struct {
	CourseRelations []*CourseRelationHTTP `json:"courseRelations"`
}

func (CourseRelationHTTPList) IsCourseRelationsResult() {}

type CoursesListHTTP struct {
	Results    []*CourseHTTP `json:"results"`
	Pagination *Pagination   `json:"pagination"`
	CountRows  int           `json:"countRows"`
}

func (CoursesListHTTP) IsCoursesResult() {}

type DeletedCourseRelation struct {
	CourseRelationID string `json:"courseRelationId"`
}

type DeletedParent struct {
	ParentID string `json:"parentId"`
}

type DeletedProjectPage struct {
	ProjectPageID string `json:"projectPageId"`
}

type DeletedRobboGroup struct {
	RobboGroupID string `json:"robboGroupId"`
}

type DeletedRobboUnit struct {
	RobboUnitID string `json:"robboUnitId"`
}

type DeletedStudent struct {
	StudentID string `json:"studentId"`
}

type DeletedTeacher struct {
	TeacherID string `json:"teacherId"`
}

type DeletedUnitAdmin struct {
	UnitAdminID string `json:"unitAdminId"`
}

type EnrollmentHTTP struct {
	Created  string `json:"created"`
	Mode     string `json:"mode"`
	IsActive bool   `json:"isActive"`
	User     string `json:"user"`
	CourseID string `json:"course_id"`
}

type EnrollmentsListHTTP struct {
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []*EnrollmentHTTP `json:"results"`
}

func (EnrollmentsListHTTP) IsEnrollmentsResult() {}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (Error) IsSignInResult() {}

func (Error) IsCohortResult() {}

func (Error) IsCourseRelationResult() {}

func (Error) IsCourseRelationsResult() {}

func (Error) IsCourseResult() {}

func (Error) IsCoursesResult() {}

func (Error) IsEnrollmentsResult() {}

func (Error) IsParentsResult() {}

func (Error) IsParentResult() {}

func (Error) IsPairsStudentParentsResult() {}

func (Error) IsProjectPageResult() {}

func (Error) IsRobboGroupResult() {}

func (Error) IsRobboGroupsResult() {}

func (Error) IsRobboUnitResult() {}

func (Error) IsRobboUnitsResult() {}

func (Error) IsStudentResult() {}

func (Error) IsStudentsResult() {}

func (Error) IsTeacherResult() {}

func (Error) IsTeachersResult() {}

func (Error) IsUnitAdminResult() {}

func (Error) IsUnitAdminsResult() {}

func (Error) IsSuperAdminResult() {}

type ImageHTTP struct {
	ID    string `json:"id"`
	Raw   string `json:"raw"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type MediaHTTP struct {
	ID  string `json:"id"`
	URI string `json:"uri"`
}

type NewAccessCourseRelationRobboGroup struct {
	CourseID     string `json:"courseId"`
	RobboGroupID string `json:"robboGroupId"`
}

type NewAccessCourseRelationRobboUnit struct {
	CourseID    string `json:"courseId"`
	RobboUnitID string `json:"robboUnitId"`
}

type NewAccessCourseRelationStudent struct {
	CourseID  string `json:"courseId"`
	StudentID string `json:"studentId"`
}

type NewAccessCourseRelationTeacher struct {
	CourseID  string `json:"courseId"`
	TeacherID string `json:"teacherId"`
}

type NewAccessCourseRelationUnitAdmin struct {
	CourseID    string `json:"courseId"`
	UnitAdminID string `json:"unitAdminId"`
}

type NewCohort struct {
	Name           string `json:"name"`
	AssignmentType string `json:"assignment_type"`
	CourseID       string `json:"course_id"`
}

type NewParent struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type NewRobboGroup struct {
	Name        string `json:"name"`
	RobboUnitID string `json:"robboUnitId"`
}

type NewRobboUnit struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type NewStudent struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	ParentID   string `json:"parentId"`
}

type NewTeacher struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type NewUnitAdmin struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type Pagination struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Count    int    `json:"count"`
	NumPages int    `json:"num_pages"`
}

type ParentHTTP struct {
	UserHTTP *UserHTTP      `json:"userHttp"`
	Children []*StudentHTTP `json:"children"`
}

func (ParentHTTP) IsParentResult() {}

func (ParentHTTP) IsGetUserResult() {}

type ParentHTTPList struct {
	Parents   []*ParentHTTP `json:"parents"`
	CountRows int           `json:"countRows"`
}

func (ParentHTTPList) IsParentsResult() {}

type ProjectPageHTTP struct {
	ProjectPageID string `json:"projectPageId"`
	LastModified  string `json:"lastModified"`
	ProjectID     string `json:"projectId"`
	Instruction   string `json:"instruction"`
	Notes         string `json:"notes"`
	Preview       string `json:"preview"`
	LinkScratch   string `json:"linkScratch"`
	Title         string `json:"title"`
	IsShared      bool   `json:"isShared"`
}

func (ProjectPageHTTP) IsProjectPageResult() {}

type ProjectPageHTTPList struct {
	ProjectPages []*ProjectPageHTTP `json:"projectPages"`
}

func (ProjectPageHTTPList) IsProjectPageResult() {}

type RobboGroupHTTP struct {
	ID           string         `json:"id"`
	LastModified string         `json:"lastModified"`
	Name         string         `json:"name"`
	RobboUnitID  string         `json:"robboUnitId"`
	Students     []*StudentHTTP `json:"students"`
}

func (RobboGroupHTTP) IsRobboGroupResult() {}

type RobboGroupHTTPList struct {
	RobboGroups []*RobboGroupHTTP `json:"robboGroups"`
	CountRows   int               `json:"countRows"`
}

func (RobboGroupHTTPList) IsRobboGroupsResult() {}

type RobboUnitHTTP struct {
	ID           string `json:"id"`
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	City         string `json:"city"`
}

func (RobboUnitHTTP) IsRobboUnitResult() {}

type RobboUnitHTTPList struct {
	RobboUnits []*RobboUnitHTTP `json:"robboUnits"`
	CountRows  int              `json:"countRows"`
}

func (RobboUnitHTTPList) IsRobboUnitsResult() {}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserRole int    `json:"userRole"`
}

type SingInResponse struct {
	AccessToken string `json:"accessToken"`
}

func (SingInResponse) IsSignInResult() {}

type StudentHTTP struct {
	UserHTTP     *UserHTTP `json:"userHttp"`
	RobboGroupID string    `json:"robboGroupId"`
	RobboUnitID  string    `json:"robboUnitId"`
}

func (StudentHTTP) IsStudentResult() {}

func (StudentHTTP) IsGetUserResult() {}

type StudentHTTPList struct {
	Students  []*StudentHTTP `json:"students"`
	CountRows int            `json:"countRows"`
}

func (StudentHTTPList) IsStudentsResult() {}

type StudentParentsHTTP struct {
	Student *StudentHTTP  `json:"student"`
	Parents []*ParentHTTP `json:"parents"`
}

type StudentParentsHTTPList struct {
	PairsStudentParents []*StudentParentsHTTP `json:"pairsStudentParents"`
}

func (StudentParentsHTTPList) IsPairsStudentParentsResult() {}

type SuperAdminHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

func (SuperAdminHTTP) IsSuperAdminResult() {}

func (SuperAdminHTTP) IsGetUserResult() {}

type TeacherHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

func (TeacherHTTP) IsTeacherResult() {}

func (TeacherHTTP) IsGetUserResult() {}

type TeacherHTTPList struct {
	Teachers  []*TeacherHTTP `json:"teachers"`
	CountRows int            `json:"countRows"`
}

func (TeacherHTTPList) IsTeachersResult() {}

type UnitAdminHTTP struct {
	UserHTTP *UserHTTP `json:"userHttp"`
}

func (UnitAdminHTTP) IsUnitAdminResult() {}

func (UnitAdminHTTP) IsGetUserResult() {}

type UnitAdminHTTPList struct {
	UnitAdmins []*UnitAdminHTTP `json:"unitAdmins"`
	CountRows  int              `json:"countRows"`
}

func (UnitAdminHTTPList) IsUnitAdminsResult() {}

type UpdateProfileInput struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type UpdateProjectPage struct {
	ProjectID     string `json:"projectId"`
	ProjectPageID string `json:"projectPageId"`
	Instruction   string `json:"instruction"`
	Notes         string `json:"notes"`
	Title         string `json:"title"`
	IsShared      bool   `json:"isShared"`
}

type UpdateRobboGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	RobboUnitID string `json:"robboUnitId"`
}

type UpdateRobboUnit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type UserHTTP struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	CreatedAt  string `json:"createdAt"`
}
