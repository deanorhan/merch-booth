-- name: ListMerch :many
select * from merch;

-- name: CreateMerch :one
INSERT INTO merch (
  vendor, title, status, price
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;