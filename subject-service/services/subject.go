package services

import (
	"subject-service/errs"
	"subject-service/repositories"
)

type SubjectService interface{}

type subjectService struct {
	subjectRepo repositories.SubjectRepository
}

func NewSubjectService(subjectRepo repositories.SubjectRepository) SubjectService {
	return subjectService{subjectRepo}
}

func (s subjectService) CreateSubject(subject repositories.Subject) error {
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

// func (s subjectService) Delete(subject repositories.Subject) error {
// 	if subject.SubjectName == "" || subject.Username == "" || subject.Day == "" || subject.Time == "" {
// 		return errs.InvalidRequest
// 	}

// 	if !checkSubjectTime(subject.Time) {
// 		return errs.InvalidRequest
// 	}

// 	if !checkSubjectDay(subject.Day) {
// 		return errs.InvalidRequest
// 	}

// 	err := s.subjectRepo.CreateSubject(subject)
// 	if err != nil {
// 		return errs.ServerError
// 	}

// 	return nil
// }
