/**
 * Trade orders submitted by the user.
**/
DROP TABLE IF EXISTS public.client_orders;
CREATE TABLE IF NOT EXISTS public.client_orders (
  id UUID PRIMARY KEY,
  symbol text,
  price NUMERIC not null,
  order_size NUMERIC not null
)
