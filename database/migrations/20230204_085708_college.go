package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type College_20230204_085708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &College_20230204_085708{}
	m.Created = "20230204_085708"

	migration.Register("College_20230204_085708", m)
}

// Run the migrations
func (m *College_20230204_085708) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE college(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`num` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *College_20230204_085708) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `college`")
}
