/*
  task 任务表
 */
CREATE TABLE `task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` text COMMENT '需要下载的链接',
  `type` int(11) DEFAULT NULL COMMENT '下载链接类型',
  `status` int(11) DEFAULT NULL COMMENT '下载状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;