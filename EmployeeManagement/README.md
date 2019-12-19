This application manages the employees of a company. It can store, list, print and delete employees.

Requirements:
- Golang (tested with go1.13.4 darwin/amd64)
- MySQL Community Server installed (tested on "macOS 10.14 (x86, 64-bit), DMG Archive" from https://dev.mysql.com/downloads/file/?id=490317)
    - install guide for MacBook is here: http://www.ccs.neu.edu/home/kathleen/classes/cs3200/MySQLMAC.pdf
- mysql is running, and it has a database employee, with user: employee and password: employeePassword1234
    - create user 'employee'@'localhost' identified by 'employeePassword1234';
    - create database employee;
    - grant all on employee.* to 'employee'@'localhost';
    - flush privileges;



This application exposes a REST API and an html page. The rest endpoints are:
    - /createEmployee  <- will add a new employee into the database, if its manager
    exists and its ID is not already there, and its fields are valid
    - /listAllEmployees <- will return a formated list of all existing employees
