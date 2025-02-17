-- Write your migrate up statements here
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

---- create above / drop below ----

DROP TABLE authors;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
