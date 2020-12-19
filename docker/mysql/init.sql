CREATE DATABASE IF NOT EXISTS Lots;
USE Lots;

-- TODO create tables

create table Provider -- Поставщики
(
	id bigint unsigned not null primary key,
	fz varchar(255) not null,
	published_at datetime not null,
	nmck varchar(255) not null,
	security_amount varchar(255) not null,
	winner varchar(255) null,
	provider_region varchar(255) null,
	address varchar(255) null,
	leader_fio varchar(255) null,
	INN varchar(255) null,
	email varchar(255) null,
	email_2 varchar(255) null,
	win_count int default 0 null,
	participate_count int default 0 null,
	last_win_date datetime null,
	purpose text null,
	procedure_type varchar(255) null,
	playground varchar(255) null,
	customer varchar(255) null,
	win_count_year int default 0 null,
	bad bool null
	activity_field  varchar(255) null,
	phone_number varchar(255) null
);

create table Purchase -- "База по закупкам"
(
	id bigint unsigned not null primary key,
	fz varchar(255) not null,
	customer varchar(255) null,
	customer_link varchar(255) null,
	customer_inn varchar(255) null,
	customer_region varchar(255) null,
	bidding_region varchar(255) null,
	customer_activity_field varchar(255) null,
	bidding_volume varchar(255) null,
	bidding_count int null,
	purchase_target varchar(255) null,
	registry_bidding_number varchar(255) null,
	contract_price varchar(255) null,
	participation_security_amount varchar(255) null,
	execution_security_amount varchar(255) null,
	published_at datetime null,
	requisition_deadline_at datetime null,
	contract_start_at datetime null,
	contract_end_at datetime null,
	playground varchar(255) null,
	purchase_link varchar(255) null
);

create table PurchaseRequisition -- "Поданные заявки"
(
	purchase_id bigint unsigned,
	provider_id bigint unsigned
);
