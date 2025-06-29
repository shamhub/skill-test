package server

import (
	"context"
)

type StudentID string
type StudentIDKey string

const studentIdKey StudentIDKey = "id"

// extract student id from
func addStudentIdToContext(ctx context.Context, id string) context.Context {
	studentId := StudentID(id)

	ctx = context.WithValue(ctx, studentIdKey, studentId)
	return ctx
}

func GetStudentId(ctx context.Context) string {
	id, ok := ctx.Value(studentIdKey).(StudentID) // creates a copy of id
	if !ok {
		panic("pass student id")
	}
	return string(id)
}

func validateRouteVariables(vars map[string]string) {
	// validate route variables for presence of valid student id otherwise panic()
}
