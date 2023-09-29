package crud_course

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrCourseNotFound = errors.New("Course not found")
