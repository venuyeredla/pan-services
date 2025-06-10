mysql -uroot -pecompwd#24

show databases;
CREATE DATABASE ecom;
use ecom;
show tables;

create table USER(id int NOT NULL AUTO_INCREMENT, first_name varchar(25), last_name varchar(25), email varchar(30), pwd varchar(25), PRIMARY KEY (id));

insert into USER(first_name, last_name, email,pwd) values ("venu","gopal","venugopal@ecom.com","ecom#24");

ALTER TABLE ecom_user RENAME COLUMN eamil to email;

