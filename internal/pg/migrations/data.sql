create table if not exists orders
(
    order_uid          varchar primary key not null,
    track_number       varchar,
    entry              varchar,
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar,
    delivery_service   varchar,
    shardkey           varchar,
    sm_id              integer,
    date_created       timestamp,
    oof_shard          varchar
);

create table if not exists deliveries
(
    order_uid varchar primary key not null references orders(order_uid) on delete cascade,
    name    varchar,
    phone   varchar,
    zip     varchar,
    city    varchar,
    address varchar,
    region  varchar,
    email   varchar
);

create table if not exists payments
(
    order_uid     varchar primary key not null references orders(order_uid) on delete cascade,
    transaction   varchar,
    request_id    varchar,
    currency      varchar,
    provider      varchar,
    amount        integer,
    payment_dt    integer,
    bank          varchar,
    delivery_cost integer,
    goods_total   integer,
    custom_fee    integer
);

create table if not exists items
(
    item_id	     serial primary key not null,
    order_uid    varchar not null references orders(order_uid) on delete cascade,
    chrt_id      integer,
    track_number varchar,
    price        integer,
    rid          varchar,
    name         varchar,
    sale         integer,
    size         varchar,
    total_price  integer,
    nm_id        integer,
    brand        varchar,
    status       integer
);