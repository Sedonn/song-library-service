-- create "artists" table
CREATE TABLE "public"."artists" (
  "id" bigserial NOT NULL,
  "name" character varying(130) NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_artists_name" to table: "artists"
CREATE UNIQUE INDEX "idx_artists_name" ON "public"."artists" ("name");
-- create "songs" table
CREATE TABLE "public"."songs" (
  "id" bigserial NOT NULL,
  "name" character varying(130) NULL,
  "artist_id" bigint NULL,
  "release_date" timestamptz NULL,
  "text" text NULL,
  "link" character varying(150) NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_songs_artist" FOREIGN KEY ("artist_id") REFERENCES "public"."artists" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create index "idx_songs_name" to table: "songs"
CREATE INDEX "idx_songs_name" ON "public"."songs" ("name");
