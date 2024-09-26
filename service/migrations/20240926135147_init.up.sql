-- create "songs" table
CREATE TABLE "public"."songs" (
  "id" bigserial NOT NULL,
  "name" character varying(130) NULL,
  "group" character varying(130) NULL,
  "release_date" character varying(10) NULL,
  "text" text NULL,
  "link" character varying(150) NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_songs_group" to table: "songs"
CREATE INDEX "idx_songs_group" ON "public"."songs" ("group");
-- create index "idx_songs_name" to table: "songs"
CREATE INDEX "idx_songs_name" ON "public"."songs" ("name");
