package models

type File struct {
	Name string
	Path string
}

type Student struct {
	Name  string
	Files []File
	OK    bool
}

func (s Student) FilenameSlice() []string {
	names := []string{}

	for _, file := range s.Files {
		names = append(names, file.Name)
	}

	return names
}

type UploadResultResponse struct {
	Students    []string  `json:"students"`
	StudentData []Student `json:"student_data"`
}
