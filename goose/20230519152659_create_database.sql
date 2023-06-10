-- +goose Up
-- +goose StatementBegin
CREATE TABLE "chat"(
"ID" serial primary key,
"moderated" smallint not null,
"name" varchar(255),
"tg" bigint not null unique);

CREATE TABLE "user"(
"admin" smallint not null,
"ID" serial primary key,
"name" varchar(255),
"tg" bigint not null);

CREATE TABLE "ban"(
"ID" serial primary key,
"chat" int not null,
"user" int not null,
"finishTime" int,
"finishUser" timestamp,
"finishReason" varchar(255),
"startTime" int not null,
"startUser" timestamp not null,
"startReason" varchar(255)  not null,
"tg" bigint not null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "ban";
DROP TABLE "chat";
DROP TABLE "user";
-- +goose StatementEnd
