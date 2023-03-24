CREATE TABLE IF NOT EXISTS "challenge"."events" (
  id BIGSERIAL PRIMARY KEY,
  "hash" TEXT UNIQUE NOT NULL,
  "block" INTEGER NOT NULL,
  "timestamp" INTEGER NOT NULL,
  "amount" TEXT NOT NULL,
  "from" TEXT NOT NULL
);