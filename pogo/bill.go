package pogo

import "github.com/jinzhu/gorm"

//go:generate go run "github.com/dirkarnez/go-dynamic-proxy/tool"

type Bill struct {
	gorm.Model
	Name        string  `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Price 		int
}

/*
CREATE TABLE `bills` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`created_at` timestamp NULL DEFAULT NULL,
`updated_at` timestamp NULL DEFAULT NULL,
`deleted_at` timestamp NULL DEFAULT NULL,
`name` varchar(255) DEFAULT NULL,
`price` int(11) DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `idx_bills_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=latin1;
*/