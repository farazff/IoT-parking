create table public.parkings
(
    id         integer default nextval('parking_id_seq'::regclass) not null
        constraint parking_pk
            primary key,
    name       varchar(256)                                        not null,
    address    varchar(256)                                        not null,
    phone      varchar(256)                                        not null,
    enabled    boolean                                             not null,
    created_at timestamp                                           not null,
    updated_at timestamp,
    deleted_at timestamp,
    uuid       uuid                                                not null
        constraint parkings_pk
            unique
);

alter table public.parkings
    owner to postgres;

create table public.zones
(
    id                integer default nextval('zone_id_seq'::regclass) not null
        constraint zone_pk
            primary key,
    capacity          integer                                          not null,
    enabled           boolean                                          not null,
    updated_at        timestamp,
    deleted_at        timestamp,
    created_at        timestamp,
    remained_capacity integer                                          not null,
    parking_id        uuid                                             not null
        constraint zones_parkings_uuid_fk
            references public.parkings (uuid)
);

alter table public.zones
    owner to postgres;

create table public.logs
(
    id         integer default nextval('parking_id_seq'::regclass) not null
        constraint logs_pk
            primary key,
    car_tag    varchar(256)                                        not null,
    enter_time timestamp                                           not null,
    exit_time  timestamp,
    parking_id uuid                                                not null
        constraint logs_parkings_uuid_fk
            references public.parkings (uuid)
);

alter table public.logs
    owner to postgres;

create table public.whitelist
(
    id         integer default nextval('parking_id_seq'::regclass) not null
        constraint whitelist_pk
            primary key,
    car_tag    varchar(256)                                        not null,
    parking_id uuid                                                not null
        constraint whitelist_parkings_uuid_fk
            references public.parkings (uuid)
);

alter table public.whitelist
    owner to postgres;

create table public.parking_admins
(
    id         integer default nextval('parking_id_seq'::regclass) not null
        constraint parking_admins_pk
            primary key,
    first_name varchar(255)                                        not null,
    last_name  varchar(255)                                        not null,
    phone      varchar(25)                                         not null,
    parking_id uuid                                                not null
        constraint parking_admins_parkings_uuid_fk
            references public.parkings (uuid),
    enabled    boolean                                             not null,
    created_at timestamp                                           not null,
    updated_at timestamp                                           not null,
    deleted_at timestamp
);

alter table public.parking_admins
    owner to postgres;

