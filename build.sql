create table if not exists users(
    id              int(11),
    name            varchar(254),
    mail            varchar(254),   # mail address
    bio             text,           # biography
    payment         int(11),        # payment price
    paymented_at    datetime,
    updated_at      datetime,
    created_at      datetime
);
create table if not exists datas(
    id              int(11),
    uid             int(11),        # db.users.id
    env             tinyint(1),
    json            varchar(254),   # file path
    updated_at      datetime,
    created_at      datetime
);