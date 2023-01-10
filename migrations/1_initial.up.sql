Begin;

-- create table users (
--     id serial primary key,
--     id_tg int not null
-- );

-- create sequence users_seq;

create table produkts (
    id serial primary key,
    id_user int not null,
    name varchar(50) not null,
    weight decimal(5,3) not null,
    bought boolean default false,
    used boolean default false,
    trown_out boolean default false,
    date_start date not null,
    date_finish date default null
);

create sequence produkts_seq;

-- alter table produkts add constraint product_users_FK
-- foreign key (id_user)
-- references users (id);

COMMIT;