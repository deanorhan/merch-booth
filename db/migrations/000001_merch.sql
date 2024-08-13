-- +goose Up
CREATE TABLE merch (
  merch_id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  vendor int NOT NULL,
  status tinyint NOT NULL,
  title text NOT NULL,
  price decimal(10, 2) NOT NULL,
  description text,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP -- ON UPDATE CURRENT_TIMESTAMP
);

-- CREATE UNIQUE INDEX vendor_to_merch_uidx ON merch (vendor, merch_id);

-- +goose Down
DROP TABLE merch;