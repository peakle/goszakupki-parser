CREATE DATABASE IF NOT EXISTS Loats;
USE Loats;

-- TODO create tables

create table Tenders
(
	fz varchar(255) not null,
	published_at datetime not null,
	nmck varchar(255) not null,
	provision_amount varchar(255) not null,
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
);
