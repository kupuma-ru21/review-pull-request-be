CREATE TABLE IF NOT EXISTS Perspectives(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Content VARCHAR (127) NOT NULL UNIQUE,
    PRIMARY KEY (ID)
)