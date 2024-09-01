
create table if not exists user_roles (
    id integer primary key autoincrement,
    user_role varchar(60),
    info varchar(250)
);

create table if not exists users (
    id integer primary key autoincrement,
    role_id integer not null,    
    username varchar(60) unique not null,
    pass_hash varchar(80) not null,
    email varchar(120) unique not null,
    foreign key (role_id) references user_roles(id)
);

-- spring, summer, fall, winter 
-- (could be more efficient to be handled as golang constants)
create table if not exists seasons (
    id integer primary key autoincrement,
    season varchar(120),
    starts date,
    ends date
);

create table if not exists holidays (
    id integer primary key autoincrement,
    season_id integer,
    foreign key (season_id) references seasons(id)
);

create table if not exists addresses (
    street varchar(200) not null,
    street2 varchar(200),
    city varchar(120),
    'state' varchar(3),
    zip_code varchar(10)
);

create table if not exists branch_locations (
    id integer primary key autoincrement,
    codename varchar(120),
    address_id integer,
    freezer_space integer,
    foreign key (address_id) references addresses(id)
);

create table if not exists vendors (
    id integer primary key autoincrement,
    vendor varchar(120) unique,
    -- order day references day of week to order
    order_day integer,
    delivery_day integer
);

create table if not exists stock_items (
    id integer primary key autoincrement,
    vendor_id integer,
    ordered_by_role integer,
    item_name varchar(100),
    ctn_amount integer, --amount in ctn
    par_amount integer, --par amount
    lifetime_sales integer,
    ytd_sales integer,
    mtd_sales integer,
    wktd_sales integer, -- week to date sales
    average_daily_sales real,
    average_weekly_sales real,
    average_monthly_sales real,
    average_yearly_sales real
);

create table if not exists stock_item_holiday_sales (
    id integer primary key autoincrement,
    stock_item_id integer,
    holiday_id integer,
    total_sales integer,
    foreign key (stock_item_id) references stock_items(id),
    foreign key (holiday_id) references holidays(id)
);

create table if not exists stock_item_quantity (
    id integer primary key autoincrement,
    branch_id integer,
    stock_item_id integer,
    -- individual quantity
    quantity integer,
    foreign key (branch_id) references branch_locations(id),
    foreign key (stock_item_id) references stock_items(id)
);

