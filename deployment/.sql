create table if not exists public.parkings
(
    id        serial not null
    constraint parking_pk
    primary key,
    name       varchar(256)                                          not null,
    address    varchar(256)                                          not null,
    phone      varchar(256)                                          not null,
    enabled    boolean                                               not null,
    created_at timestamp default now()                               not null,
    updated_at timestamp default now(),
    deleted_at timestamp,
    uuid       uuid
    );

alter table public.parkings
    owner to postgres;

create table if not exists public.parking_admins
(
    id         serial not null
    constraint parking_admins_pk
    primary key,
    first_name varchar(255)                                          not null,
    last_name  varchar(255)                                          not null,
    phone      varchar(25)                                           not null,
    enabled    boolean                                               not null,
    created_at timestamp default now()                               not null,
    updated_at timestamp default now()                               not null,
    deleted_at timestamp,
    password   varchar(255),
    parking_id integer                                               not null
    constraint parking_admins_parkings_id_fk
    references public.parkings
    );

alter table public.parking_admins
    owner to postgres;



create table if not exists public.whitelists
(
    id         serial not null
    constraint whitelist_pk
    primary key,
    car_tag    varchar(256)                                        not null,
    parking_id integer                                             not null
    constraint whitelists_parkings_id_fk
    references public.parkings
    );

alter table public.whitelists
    owner to postgres;

create table if not exists public.logs
(
    id         serial not null
    constraint logs_pk
    primary key,
    car_tag    varchar(256)                                        not null,
    enter_time timestamp                                           not null,
    exit_time  timestamp,
    parking_id integer                                             not null
    constraint logs_parkings_id_fk
    references public.parkings
    );

alter table public.logs
    owner to postgres;



create table if not exists public.zones
(
    id                serial not null
    constraint zone_pk
    primary key,
    capacity          integer                                          not null,
    enabled           boolean                                          not null,
    updated_at        timestamp,
    deleted_at        timestamp,
    created_at        timestamp,
    remained_capacity integer                                          not null,
    parking_id        integer                                          not null
    constraint zones_parkings_id_fk
    references public.parkings
);

alter table public.zones
    owner to postgres;




