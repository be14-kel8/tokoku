show databases;
use tom14;

-- DROP 
DROP TABLE user_activity;
DROP TABLE fotos;
DROP TABLE activities;
DROP TABLE users;

CREATE TABLE users (
	user_id int auto_increment,
	user_name varchar(255) DEFAULT NULL,
	password varchar(255) DEFAULT NULL,
	primary key (user_id)
);

CREATE TABLE activities (
	act_id int auto_increment,
	title varchar(255) NOT NULL,
	location varchar(255) DEFAULT NULL,
	create_date datetime NOT NULL,
	owner int NOT NULL,
	primary key (act_id),
	CONSTRAINT fk_activities_users_tom
	FOREIGN KEY (owner) references users(user_id)
);

CREATE TABLE fotos (
	id int,
	foto varchar(255),
	PRIMARY KEY (id),
	CONSTRAINT fk_fotos_activities
	FOREIGN KEY (id) references activities(act_id)
);

CREATE TABLE `user_activity` (
	id_user int NOT NULL,
	id_activity int NOT NULL,
	execute_date timestamp DEFAULT now(),
	CONSTRAINT fk_user_activity_users FOREIGN KEY (id_user) REFERENCES users(user_id),
	CONSTRAINT fk_user_activity_activities FOREIGN KEY (id_activity) REFERENCES activities(act_id),
	PRIMARY KEY (id_user, id_activity)
);


-- INSERT 

insert into users 
values (1,"Thomas","thomasoke");

insert into users(user_name,password)
values	("Gianto","patenkali")



