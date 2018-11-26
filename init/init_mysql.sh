#!/bin/bash
MYSQL="mysql -u -h -p "

#创建数据库
function create_database(){
    ${MYSQL} < sql/create_database.sql
}


#创建表
function create_table (){

    ${MYSQL} < sql/create_table.sql
}