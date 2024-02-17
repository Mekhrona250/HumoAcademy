package service

import (
	"fmt"
	"humoAcademy/pkg/errors"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) GetListOfUsersByCourseID(userID, courseID int) (xlsx *excelize.File, err error) {

	user, err := s.Repo.GetUserByID(userID)

	if err != nil {
		return
	}

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}

	applications, err := s.Repo.GetAllApplicationsByCourseID(courseID)

	if err != nil {
		return
	}

	xlsx = excelize.NewFile()

	xlsx.SetCellValue("Sheet1", "A1", "Name")
	xlsx.SetCellValue("Sheet1", "B1", "Surname")
	xlsx.SetCellValue("Sheet1", "C1", "Date of birth")
	xlsx.SetCellValue("Sheet1", "D1", "Phone number")

	for i, v := range applications {
		user, tempErr := s.Repo.GetUserByID(v.UserId)

		if tempErr != nil {
			err = errors.ErrDataNotFound
		}

		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+2), user.Name)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+2), user.Surname)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%v", i+2), user.DateOfBirth.Format("02.01.2006"))
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("D%v", i+2), user.PhoneNumber)
	}

	return

}
