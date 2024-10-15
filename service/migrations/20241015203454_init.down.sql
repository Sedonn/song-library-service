-- reverse: create index "idx_songs_name" to table: "songs"
DROP INDEX "public"."idx_songs_name";
-- reverse: create "songs" table
DROP TABLE "public"."songs";
-- reverse: create index "idx_artists_name" to table: "artists"
DROP INDEX "public"."idx_artists_name";
-- reverse: create "artists" table
DROP TABLE "public"."artists";
