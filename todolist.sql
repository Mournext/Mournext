create table `category`
    -- --------------------------------------------------
    --  Table Structure for `todolist/models.Category`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `category` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `task_count` bigint NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

create table `task`
    -- --------------------------------------------------
    --  Table Structure for `todolist/models.Task`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `task` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `cate` varchar(255) NOT NULL DEFAULT '' ,
        `content` varchar(5000) NOT NULL DEFAULT '' ,
        `attachment` varchar(255) NOT NULL DEFAULT '' ,
        `created_date` datetime NOT NULL,
        `last_modified_at` datetime NOT NULL,
        `finish_date` datetime NOT NULL
    ) ENGINE=InnoDB;
    CREATE INDEX `task_created_date` ON `task` (`created_date`);
    CREATE INDEX `task_last_modified_at` ON `task` (`last_modified_at`);
    CREATE INDEX `task_finish_date` ON `task` (`finish_date`);
