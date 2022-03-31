USE db

CREATE TABLE IF NOT EXISTS companies
(
	id int auto_increment primary key,
    name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS phones
(
	id int auto_increment primary key,
	company_id int not null,
	number varchar(30), foreign key (company_id) REFERENCES companies(id)
);