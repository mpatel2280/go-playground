-- Classes
-- name: CreateClass :exec
INSERT INTO classes (name) VALUES (?);

-- name: ListClasses :many
SELECT * FROM classes;

-- Students
-- name: CreateStudent :exec
INSERT INTO students (name, class_id) VALUES (?, ?);

-- name: GetStudentWithClass :one
SELECT s.id, s.name, c.id AS class_id, c.name AS class_name
FROM students s
LEFT JOIN classes c ON s.class_id = c.id
WHERE s.id = ?;

-- name: ListStudents :many
SELECT * FROM students;

-- Assignments
-- name: CreateAssignment :exec
INSERT INTO assignments (title, student_id) VALUES (?, ?);

-- name: ListAssignmentsByStudent :many
SELECT * FROM assignments WHERE student_id = ?;
