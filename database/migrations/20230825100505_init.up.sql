CREATE TABLE public.user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25),
    password VARCHAR(255),
    email VARCHAR(50),
    telephone VARCHAR(13),
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
   
);

CREATE TABLE public.wallet (
    id SERIAL PRIMARY KEY, 
    userId INT,
    balance DECIMAL(22, 1),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
   
) ;

CREATE TABLE public.wallet_log (
    id SERIAL PRIMARY KEY, 
    userId INT,
    walletId INT,
    walletIn DECIMAL(22, 1),
    walletOut DECIMAL(22, 1),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
   
) ;

CREATE TABLE public.bank_account (
    id SERIAL PRIMARY KEY, 
    userId INT,
    accountNumber VARCHAR(50),
    accountName VARCHAR(50),
    bankName VARCHAR(50),
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
   
);