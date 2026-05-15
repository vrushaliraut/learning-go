package main

import (
	"errors"
	"fmt"
	"time"
)

// School system -
// 1 school - M classes
// 1 class - M students
// 1 student --> many courses
// 1 course -> 1 teacher

// Actions -
/*
## Core CRUD Operations

**School**
- Create/close a school
- Update school details (name, address, accreditation)
- List all classes in a school
- Get school statistics (total students, teachers, classes)

**Class**
- Create/disband a class
- Assign students to a class
- Remove students from a class
- Move student between classes
- Get class roster
- Get class schedule (all courses)

**Student**
- Enroll/drop a student
- Update student information
- Enroll in a course
- Drop a course
- View student's transcript/grades
- View student's current course load
- Check graduation requirements progress

**Course**
- Create/modify a course
- Assign a teacher to a course
- Change teacher for a course
- Get course roster (all enrolled students)
- Get course capacity/waitlist

**Teacher**
- Hire/terminate a teacher
- Update teacher information
- View teacher's course load
- View teacher's schedule

## Relationship Behaviors

**Enrollment Constraints**
- Prevent enrolling in conflicting time slots
- Enforce prerequisites (Course A before Course B)
- Enforce maximum class size
- Handle waitlists when course is full
- Prevent duplicate enrollments

**Data Integrity**
- Cascade delete: removing a school removes all classes/students/courses
- Restrict delete: prevent deleting a teacher who has active courses
- Archive vs. hard delete (keep historical records)

## Query/Reporting Behaviors

- Search students by name, ID, or class
- Filter courses by teacher, subject, or availability
- Generate report cards
- Calculate GPA/rankings
- Identify students with course conflicts
- Find teachers with availability
- Export class rosters

## Validation Behaviors

- Ensure a student isn't in multiple classes simultaneously (if classes are time-based)
- Validate teacher doesn't exceed maximum course load
- Check student isn't enrolled in the same course twice
- Verify course exists before enrollment

## Administrative Behaviors

- Academic year rollover (promote students, archive old data)
- Bulk import students/teachers
- Generate enrollment reports
- Handle transfers (student moves between schools)

Would you like me to elaborate on any specific behavior or discuss how to model these
in a particular programming paradigm (OOP, functional, database schema)?

Entities, Object and fild behaviour in some program and try to do it in go
*/

type Creater interface {
	Create(name string) (any, error)
}
type School struct {
	name     string
	students []Student
	courses  []Course
	classes  []Class
	teachers []Teacher
}

type Statisfier interface {
	createStats() string
}

func (s *School) createStats() string {
	return fmt.Sprintf("School %s has %d students. %d teachers, %d courses,  %d classes")

}

func (s *School) Create(name string) (any, error) {
	if name == "" {
		return School{}, errors.New("name is required")
	}

	return School{name: name}, nil
}

type Class struct {
	name     string
	students []Student
}

type Student struct {
	name    string
	courses []Course
}

type Course struct {
	name     string
	duration time.Duration
}

func (c Course) createCourse(name string, duration time.Duration) Course {
	return Course{name: name, duration: duration}
}

type Teacher struct {
	Course
	name string
}

func main() {
	s := School{
		name:     "My School",
		students: nil,
		courses:  nil,
		classes:  nil,
		teachers: nil,
	}

	courses := []string{"language", "maths", "science"}

	for _, course := range courses {
		Course{name: course}.createCourse(course, time.Duration(0))
	}

	fmt.Println(s.students)
}
