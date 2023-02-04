package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Student_20230203_143355 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Student_20230203_143355{}
	m.Created = "20230203_143355"

	migration.Register("Student_20230203_143355", m)
}

// Run the migrations
func (m *Student_20230203_143355) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE student(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`code` varchar(128) NOT NULL,`age` int(11) DEFAULT NULL,`gender` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Student_20230203_143355) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `student`")
}
