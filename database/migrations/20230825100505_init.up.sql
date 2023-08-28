CREATE TABLE user(
    id INT NOT NULL AUTO_INCREMENT, 
    username VARCHAR(25),
    password VARCHAR(50),
    email VARCHAR(50),
    telephone VARCHAR(13),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) Engine = Innodb;

CREATE TABLE wallet(
    id INT NOT NULL AUTO_INCREMENT, 
    userId INT,
    balance DECIMAL(22, 1),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) Engine = Innodb;

CREATE TABLE wallet_log(
    id INT NOT NULL AUTO_INCREMENT, 
    userId INT,
    walletId INT,
    walletIn DECIMAL(22, 1),
    walletOut DECIMAL(22, 1),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) Engine = Innodb;

CREATE TABLE bank_account(
    id INT NOT NULL AUTO_INCREMENT, 
    userId INT,
    accountNumber VARCHAR(50),
    accountName VARCHAR(50),
    bankName VARCHAR(50),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) Engine = Innodb;

CREATE TABLE session(
    id INT NOT NULL AUTO_INCREMENT, 
    userId INT,
    session VARCHAR(255),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) Engine = Innodb;