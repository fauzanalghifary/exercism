package school

import "sort"

// Grade represents a grade level with its students
type Grade struct {
	Level    int
	Students []string
}

// School represents a school roster
type School struct {
	roster map[int]map[string]bool
}

func New() *School {
	return &School{
		roster: make(map[int]map[string]bool),
	}
}

func (s *School) Add(student string, grade int) {
	if s.roster[grade] == nil {
		s.roster[grade] = make(map[string]bool)
	}
	s.roster[grade][student] = true
}

func (s *School) Grade(level int) []string {
	students := make([]string, 0)
	for student := range s.roster[level] {
		students = append(students, student)
	}
	sort.Strings(students)
	return students
}

func (s *School) Enrollment() []Grade {
	grades := make([]int, 0)
	for grade := range s.roster {
		grades = append(grades, grade)
	}
	sort.Ints(grades)

	enrollment := make([]Grade, 0)
	for _, grade := range grades {
		students := s.Grade(grade)
		enrollment = append(enrollment, Grade{
			Level:    grade,
			Students: students,
		})
	}
	return enrollment
}
