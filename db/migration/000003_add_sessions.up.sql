CREATE TABLE "sessions"(
  "id" uuid PRIMARY KEY,
  "username"  varchar NOT NULL,
  "refresh_token" varchar NOT NULL ,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
alter table "sessions" add foreign key ("username") references "users"("username");