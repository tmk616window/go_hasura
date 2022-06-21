comment on column "public"."users"."fullname" is E'users';
alter table "public"."users" alter column "fullname" drop not null;
alter table "public"."users" add column "fullname" text;
