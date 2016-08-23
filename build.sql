create table if not exists users(
    id              int(11)         NOT NULL  AUTO_INCREMENT,
    name            varchar(254)    NOT NULL,
    mail            varchar(254)    NOT NULL,                   # mail address
    bio             text,                                       # biography
    payment         int(11)         DEFAULT 0,                  # payment price
    paymented_at    datetime        DEFAULT NULL,
    updated_at      datetime        NOT NULL,
    created_at      datetime        NOT NULL
);
create table if not exists datas(
    id              int(11)         NOT NULL AUTO_INCREMENT,
    uid             int(11)         NOT NULL,                   # db.users.id
    env             tinyint(1)      NOT NULL,
    json            varchar(254),                               # file path
    updated_at      datetime        NOT NULL,
    created_at      datetime        NOT NULL
);