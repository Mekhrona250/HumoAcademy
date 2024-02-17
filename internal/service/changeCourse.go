package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
)

func (s *Service) ChangeCourse(course models.Course, courseID int, userID int) (err error) {
	user, err := s.Repo.GetUserByID(userID)

	if err != nil {
		return
	}

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}

	oldCourse, err := s.Repo.GetCourseByID(courseID)
	if err != nil {
		return err
	}

	if course.Name != "" {
		oldCourse.Name = course.Name
	}

	if !course.StartDate.IsZero() {
		oldCourse.StartDate = course.StartDate
	}

	if course.Duration != 0 {
		oldCourse.Duration = course.Duration
	}

	if course.Schedule != "" {
		oldCourse.Schedule = course.Schedule
	}

	if course.AgeLimit != 0 {
		oldCourse.AgeLimit = course.AgeLimit
	}

	if !course.RegistrationEndDate.IsZero() {
		oldCourse.RegistrationEndDate = course.RegistrationEndDate
	}

	if course.Address != "" {
		oldCourse.Address = course.Address
	}

	if course.Description != "" {
		oldCourse.Description = course.Description
	}

	if course.Mentor != "" {
		oldCourse.Mentor = course.Mentor
	}

	if course.Format != "" {
		oldCourse.Format = course.Format
	}

	if course.Language != "" {
		oldCourse.Language = course.Language
	}

	err = s.Repo.ChangeCourse(oldCourse, courseID)

	if err != nil {
		err = errors.ErrBadRequest
		return
	}

	return
}
