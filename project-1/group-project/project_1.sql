CREATE DATABASE project;

USE project;

CREATE TABLE users(
id INT PRIMARY KEY NOT NULL auto_increment,
nama VARCHAR(100),
gender VARCHAR(50),
telp VARCHAR(13),
password VARCHAR(20),
saldo INT(13),
created_at DATETIME default CURRENT_TIMESTAMP,
updated_at DATETIME default CURRENT_TIMESTAMP
);

CREATE TABLE status(
id INT PRIMARY KEY NOT NULL,
status VARCHAR(20),
CONSTRAINT FK_user_status FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE transfer_detail(
id INT PRIMARY KEY NOT NULL auto_increment,
user_id INT,
telp_sender VARCHAR(13),
telp_receiver VARCHAR(13),
saldo_transfer INT(13) NOT NULL,
created_at DATETIME default CURRENT_TIMESTAMP,
CONSTRAINT FK_transfer_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE topup_detail(
id INT PRIMARY KEY NOT NULL auto_increment,
user_id INT,
telp VARCHAR(13),
saldo_topup INT(13) NOT NULL,
created_at DATETIME default CURRENT_TIMESTAMP,
CONSTRAINT FK_topup_user FOREIGN KEY (user_id) REFERENCES users(id)
);


-- coba database --
INSERT INTO users (nama, gender, telp, password)
VALUES ("Yuki", "Perempuan", "0797869595", "0978@ti");

DROP DATABASE project;







