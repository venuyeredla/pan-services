docker exec -it local-mysql /bin/bash
mysql -uroot -pecompwd#24

show databases;
CREATE DATABASE ecom;
use ecom;
show tables;

create table ROLE(ROLE_ID int NOT NULL AUTO_INCREMENT,ROLE_NAME varchar(25), ROLE_DESC varchar(50), CREATE_DATE DATETIME, MODIFIED_DATE DATETIME, ACTIVE boolean, PRIMARY KEY (ROLE_ID));

insert into ROLE(ROLE_NAME,ROLE_DESC,CREATE_DATE,MODIFIED_DATE,ACTIVE) values ("ADMIN","Admistrator",NOW(), NOW(), 1);
insert into ROLE(ROLE_NAME,ROLE_DESC,CREATE_DATE,MODIFIED_DATE,ACTIVE) values ("SELLER","Seller",NOW(), NOW(), 1);
insert into ROLE(ROLE_NAME,ROLE_DESC,CREATE_DATE,MODIFIED_DATE,ACTIVE) values ("BUYER","Buyer",NOW(), NOW(), 1);

create table USER(user_id int NOT NULL AUTO_INCREMENT, first_name varchar(25), last_name varchar(25), email varchar(30), pwd varchar(25), CREATE_DATE DATETIME, MODIFIED_DATE DATETIME, ACTIVE boolean, PRIMARY KEY (user_id));


ALTER TABLE USER RENAME COLUMN id to user_id;

ALTER TABLE USER ADD COLUMN CREATE_DATE DATETIME, MODIFIED_DATE DATETIME, ACTIVE boolean;

insert into USER(first_name, last_name, email,pwd) values ("venu","gopal","venugopal@ecom.com","ecom#24");

update USER SET CREATE_DATE=NOW(), MODIFIED_DATE=NOW(), ACTIVE=1 where user_id=1;

create table USER_ROLE(
        USER_ROLE_ID int NOT NULL AUTO_INCREMENT, 
        USER_ID int NOT NULL,
        ROLE_ID int NOT NULL, 
        CREATE_DATE DATETIME, 
        MODIFIED_DATE DATETIME, 
        ACTIVE boolean, 
        PRIMARY KEY (USER_ROLE_ID),
        FOREIGN KEY (USER_ID) REFERENCES USER(user_id),
        FOREIGN KEY (ROLE_ID) REFERENCES ROLE(ROLE_ID)
        );


insert into USER_ROLE(USER_ID,ROLE_ID,CREATE_DATE,MODIFIED_DATE,ACTIVE) values(1,1,NOW(),NOW(),1);
insert into USER_ROLE(USER_ID,ROLE_ID,CREATE_DATE,MODIFIED_DATE,ACTIVE) values(1,2,NOW(),NOW(),1);
insert into USER_ROLE(USER_ID,ROLE_ID,CREATE_DATE,MODIFIED_DATE,ACTIVE) values(1,3,NOW(),NOW(),1);


select U.user_id, U.first_name, U.last_name, U.email, R.ROLE_NAME from USER U INNER JOIN USER_ROLE UR on U.user_id=UR.USER_ID INNER JOIN ROLE R on UR.ROLE_ID=R.ROLE_ID;

SELECT   U.user_id, U.first_name, U.last_name, U.email, GROUP_CONCAT(R.ROLE_NAME ORDER BY R.ROLE_NAME ASC SEPARATOR ', ') as ROLES
from USER U INNER JOIN USER_ROLE UR on U.user_id=UR.USER_ID INNER JOIN ROLE R on UR.ROLE_ID=R.ROLE_ID where U.email ='venugopal@ecom.com' AND U.pwd='ecom#24'
GROUP BY U.user_id, U.first_name, U.last_name, U.email;