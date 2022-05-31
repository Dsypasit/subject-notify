package services

import (
	"subject-service/errs"
	"subject-service/repositories"
)

// type SubjectService interface {
// 	CreateSubject(subject repositories.Subject) error
// 	DeleteSubject(subQuery repositories.SubjectQuery) error
// 	UpdateSubject(subQuery repositories.SubjectQuery, subUpdate repositories.SubjectUpdate) error
// 	GetSubjectsByUsername(subQuery repositories.SubjectQuery) ([]repositories.Subject, error)
// 	GetSubjectsBySubjectname(subQuery repositories.SubjectQuery) ([]repositories.Subject, error)
// 	GetSubjectsByQuery(subQuery repositories.SubjectQuery) ([]repositories.Subject, error)
// 	GetSubjectsByDay(subQuery repositories.SubjectQuery) ([]repositories.Subject, error)
// 	GetSubjectsByTime(subQuery repositories.SubjectQuery) ([]repositories.Subject, error)
// }

type SubjectService struct {
	subjectRepo repositories.SubjectRepository
}

func NewSubjectService(subjectRepo repositories.SubjectRepository) SubjectService {
	return SubjectService{subjectRepo}
}

func (s SubjectService) CreateSubject(subject repositories.Subject) error {
	if subject.SubjectName == "" || subject.Username == "" || subject.Day == "" || subject.Time == "" {
		return errs.InvalidRequest
	}

	if !checkSubjectTime(subject.Time) {
		return errs.InvalidRequest
	}

	if !checkSubjectDay(subject.Day) {
		return errs.InvalidRequest
	}

	err := s.subjectRepo.CreateSubject(subject)
	if err != nil {
		return errs.ServerError
	}

	return nil
}

func (s SubjectService) DeleteSubject(subQuery repositories.SubjectQuery) error {
	if subQuery.Username == "" || subQuery.SubjectName == "" {
		return errs.InvalidRequest
	}

	if subQuery.Time != "" && !checkSubjectTime(subQuery.Time) {
		return errs.InvalidRequest
	}

	if subQuery.Day != "" && !checkSubjectDay(subQuery.Day) {
		return errs.InvalidRequest
	}

	err := s.subjectRepo.DeleteSubject(subQuery)
	if err != nil {
		return errs.ServerError
	}

	return nil
}

func (s SubjectService) UpdateSubject(subQuery repositories.SubjectQuery, subUpdate repositories.SubjectUpdate) error {
	if subQuery.Username == "" || subQuery.SubjectName == "" {
		return errs.InvalidRequest
	}

	if subQuery.Time != "" && !checkSubjectTime(subQuery.Time) {
		return errs.InvalidRequest
	}

	if subQuery.Day != "" && !checkSubjectDay(subQuery.Day) {
		return errs.InvalidRequest
	}

	err := s.subjectRepo.UpdateSubjectByQuery(subQuery, subUpdate)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func (s SubjectService) GetSubjectsByUsername(subQuery repositories.SubjectQuery) ([]repositories.Subject, error) {
	if subQuery.Username == "" {
		return nil, errs.InvalidRequest
	}

	subjects, err := s.subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		return nil, errs.ServerError
	}

	return subjects, nil
}

func (s SubjectService) GetSubjectsBySubjectname(subQuery repositories.SubjectQuery) ([]repositories.Subject, error) {
	if subQuery.Username == "" || subQuery.SubjectName == "" {
		return nil, errs.InvalidRequest
	}

	subjects, err := s.subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		return nil, errs.ServerError
	}

	return subjects, nil
}

func (s SubjectService) GetSubjectsByTime(subQuery repositories.SubjectQuery) ([]repositories.Subject, error) {
	if subQuery.Username == "" || subQuery.Time == "" {
		return nil, errs.InvalidRequest
	}

	if !checkSubjectTime(subQuery.Time) {
		return nil, errs.InvalidRequest
	}

	subjects, err := s.subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		return nil, errs.ServerError
	}

	return subjects, nil
}

func (s SubjectService) GetSubjectsByDay(subQuery repositories.SubjectQuery) ([]repositories.Subject, error) {
	if subQuery.Username == "" || subQuery.Day == "" {
		return nil, errs.InvalidRequest
	}

	if !checkSubjectDay(subQuery.Day) {
		return nil, errs.InvalidRequest
	}

	subjects, err := s.subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		return nil, errs.ServerError
	}

	return subjects, nil
}

func (s SubjectService) GetSubjectsByQuery(subQuery repositories.SubjectQuery) ([]repositories.Subject, error) {
	if subQuery.Username == "" {
		return nil, errs.InvalidRequest
	}

	if subQuery.Time != "" && !checkSubjectTime(subQuery.Time) {
		return nil, errs.InvalidRequest
	}

	if subQuery.Day != "" && !checkSubjectDay(subQuery.Day) {
		return nil, errs.InvalidRequest
	}

	subjects, err := s.subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		return nil, errs.ServerError
	}

	return subjects, nil
}
