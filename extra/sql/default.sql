CREATE TABLE kingquiz.users (
	id varchar(40) NOT NULL,
	email varchar(320) NOT NULL,
	username varchar(100) NULL,
	password varchar(255) NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;