-- name: GetUser :many
SELECT * FROM public.user;

-- name: GetUserById :one
SELECT * FROM public.user WHERE id = $1;

-- name: CreateUser :one
INSERT INTO public.user (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE public.user SET name = $2, email = $3, password = $4 WHERE id = $1 RETURNING *;

-- name: UpdateUserPassword :one
UPDATE public.user SET password = $2 WHERE id = $1 RETURNING *;

-- name: UpdateUserName :one
UPDATE public.user SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM public.user WHERE id = $1 RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM public.user WHERE email = $1;

-- name: GetUserCount :one
SELECT COUNT(*) FROM public.user;