package repositories

import "gorm.io/gorm"

type Subject struct {
	ID          uint   `json:"id"`
	SubjectName string `json:"subject_name"`
	Teacher     string `json:"teacher"`
	Day         string `json:"day"`
	Time        string `json:"time"`
	Active      bool   `json:"active" gorm:"default:true"`
	Link        string `json:"link"`
	Username    string `json:"username"`
}

type SubjectQuery struct {
	Username    string `json:"username"`
	SubjectName string `json:"subject_name"`
	Day         string `json:"day"`
	Time        string `json:"time"`
}

type SubjectUpdate struct {
	SubjectName string `json:"subject_name"`
	Teacher     string `json:"teacher"`
	Day         string `json:"day"`
	Time        string `json:"time"`
	Active      bool   `json:"active"`
	Link        string `json:"link"`
}

type SubjectRepository interface {
	CreateSubject(subject Subject) error
	DeleteSubject(subQuery SubjectQuery) error
	GetSubjectByQuery(subQuery SubjectQuery) (subjects []Subject, err error)
	UpdateSubjectByQuery(subQuery SubjectQuery, subUpdate SubjectUpdate) (err error)
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) subjectRepository {
	db.AutoMigrate(&Subject{})
	return subjectRepository{db}
}

func (r subjectRepository) CreateSubject(subject Subject) error {
	return r.db.Create(&subject).Error
}

func (r subjectRepository) DeleteSubject(subQuery SubjectQuery) error {
	subject := queryTosubject(subQuery)
	return r.db.Where(&subject).Delete(&Subject{}).Error
}

func (r subjectRepository) GetSubjectByQuery(subQuery SubjectQuery) (subjects []Subject, err error) {
	subject := queryTosubject(subQuery)
	err = r.db.Where(&subject).Find(&subjects).Error
	return subjects, err
}

func (r subjectRepository) UpdateSubjectByQuery(subQuery SubjectQuery, subUpdate SubjectUpdate) (err error) {
	subjectQuery := queryTosubject(subQuery)
	subjectUpdate := updateTosubject(subUpdate)
	err = r.db.Model(&Subject{}).Where(&subjectQuery).Updates(subjectUpdate).Error
	return err
}

func queryTosubject(subQuery SubjectQuery) Subject {
	subject := Subject{
		Username:    subQuery.Username,
		SubjectName: subQuery.SubjectName,
		Day:         subQuery.Day,
		Time:        subQuery.Time,
	}
	return subject
}

func updateTosubject(subUpdate SubjectUpdate) Subject {
	subject := Subject{
		SubjectName: subUpdate.SubjectName,
		Day:         subUpdate.Day,
		Time:        subUpdate.Time,
		Teacher:     subUpdate.Teacher,
		Link:        subUpdate.Link,
	}
	return subject
}
