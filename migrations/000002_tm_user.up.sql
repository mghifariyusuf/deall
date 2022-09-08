create table IF NOT EXISTS tm_user
(
    id           varchar(255)                          not null,
    username     varchar(255)                          not null,
    email        varchar(255)                          not null,
    password     varchar(255)                          not null,
    first_name   varchar(255)                          not null,
    last_name    varchar(255)                          null,
    phone_number varchar(255)                          null,
    role_id      varchar(255)                          not null,
    created_at   timestamp default current_timestamp() not null on update current_timestamp(),
    created_by   varchar(255)                          null,
    updated_at   timestamp                             null,
    updated_by   varchar(255)                          null,
    deleted_at   timestamp                             null,
    deleted_by   varchar(255)                          null,
    CONSTRAINT tm_user_primary PRIMARY KEY (ID),
    constraint tm_user_email_uindex
        unique (email),
    constraint tm_user_id_uindex
        unique (id),
    constraint tm_user_username_uindex
        unique (username)
);

INSERT INTO tm_user 
(id, username, email, password, first_name, last_name, phone_number, role_id, created_at, created_by)
VALUES
('8e0f4c66-9782-4e09-a311-6e662d11222d', 'admin', 'admin@mail.com', '$2a$04$BpZKxYGVX4ATh3NZ4Z8IZ.Ve0OHnUeXExSOGTt4xt9p7hMraamlwO', 'admin', 'admin', '081234569', 'admin', NOW(), 'admin'),
('20744457-a850-49b2-b2dc-16c491328aff', 'user1', 'user1@mail.com', '$2a$04$fWrkF9g4FUlpa98FMdNp5O2BlKnQiNAzmbqILOZOIoSPw6.cuo2oC', 'user1', 'user1', '081234569', 'user', NOW(), 'admin')
ON DUPLICATE KEY UPDATE
  updated_at = NOW();