CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    userId VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    verfiedEmail BOOLEAN NOT NULL DEFAULT false,
    INDEX idx_userId (userId),
    INDEX idx_email (email)
);

CREATE TABLE IF NOT EXISTS recommendations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    productId VARCHAR(255) NOT NULL
);

INSERT INTO recommendations(productId) VALUES(1);
INSERT INTO recommendations(productId) VALUES(2);
INSERT INTO recommendations(productId) VALUES(3);
INSERT INTO recommendations(productId) VALUES(4);
INSERT INTO recommendations(productId) VALUES(5);
INSERT INTO recommendations(productId) VALUES(6);