CREATE TABLE `Employee` (
  `Id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '名字',
  `Salary` int NOT NULL DEFAULT 0 COMMENT '工资',
  `DepartmentId` int NOT NULL DEFAULT 0 COMMENT '部门',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 COMMENT='员工信息';

CREATE TABLE `Department` (
  `Id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `Name` varchar(128) NOT NULL DEFAULT '' COMMENT '名字',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 COMMENT='部门信息';
