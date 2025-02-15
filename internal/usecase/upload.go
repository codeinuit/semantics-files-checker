package usecase

import (
	"archive/zip"
	"path/filepath"
	"strings"

	"github.com/codeinuit/semantics-files-checker/internal/models"
	"github.com/sirupsen/logrus"
)

func CheckZipFilesSemantics(log *logrus.Logger, files []*zip.File) []models.Student {
	students := make(map[string]models.Student)

	for _, file := range files {
		slice := strings.Split(file.Name, "/")
		if len(slice) <= 1 {
			continue
		}

		student := slice[1]

		_, ok := students[student]
		if !ok {
			log.Debugf("added path: %s", student)
			students[student] = models.Student{
				Name: student,
			}
		}

		if !file.FileInfo().IsDir() {
			folder := filepath.Base(filepath.Dir(file.Name))
			log.Infof("file %s in folder %s", filepath.Base(file.Name), folder)
			stu := students[student]
			filename := filepath.Base(file.Name)
			if filepath.Ext(filename) == ".txt" {
				stu.OK = true
			}

			stu.Files = append(stu.Files, models.File{
				Name: filename,
				Path: file.Name,
			})

			students[student] = stu
		}
	}

	res := []models.Student{}
	for _, n := range students {
		res = append(res, n)
	}

	return res
}
