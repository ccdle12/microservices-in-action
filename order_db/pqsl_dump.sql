--
-- Table for an "orders"
--

DROP TABLE IF EXISTS "order";

CREATE TABLE "order" (
  order_id text UNIQUE NOT NULL,
  user_id text,
  symbol text,
  amount integer,
  status integer,
  PRIMARY KEY("order_id")
)
