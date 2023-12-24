USE `language`;
CREATE TABLE `languages`(
    `id` int Primary Key,
    `name` VARCHAR(255)
);
INSERT INTO `languages`(`id`, `name`)
VALUES(1, 'Java'),(2, 'python'),(3, 'golang');