USE `language`;
CREATE TABLE `languages`(
    `id` int Primary Key NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255)
);
INSERT INTO `languages`(`id`, `name`)
VALUES(1, 'Java'),(2, 'python'),(3, 'golang');
\! echo "Done with SQL script ############";