CREATE Database News;

use News;

Create table NewsItem(
    ID int PRIMARY KEY AUTO_INCREMENT,
    Title varchar(255),
    Description varchar(255),
    Body text,
    Type varchar(255),
    Link varchar(255),
    Source varchar(255),
    UpdatedAt datetime
);