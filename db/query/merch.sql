-- name: ListMerch :many
select * from merch;

-- name: CreateMerch :one
INSERT INTO merch (
  vendor, title, status, price
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: GetMerchItem :one
select * from merch where merch_id = ?;