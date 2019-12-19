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


Algorithm to format the data O(NlogN):
We arrange the data in a tree, where the CEO is the root, its immediate manages are its children nodes,
and so on. After building* this tree we iterate on it and compose the final result*.

*To build the tree efficiently, we first compute, for each employee, on what level in the tree
it will be placed. To do this we use an auxiliarry data structure, namely a slice of parents.

*To compute the final result we traverse the tree, starting with its lowest rigth node. The height of the
node will be equal to how many empty tabs will be in front of that employee name, and the distance to
the right will determine its index in the result list:
    - the rightmost node is at the bottom of the result list
    - we first traverse the right subtree, then its left subtree then its root
    - this traversal produces the lowest line in the result list, then the second lowest, etc



NOTE 1: we reject the entries that have an invalid manager ID, or with an already existing ID. 
This is solely for flavor, the algorithm to arrange the data could, with minimal changes, handle these as well.

NOTE 2: I went for a complexity of O(NlogN). With a modified version of the union find/disjoint set algorithm,
I believe an even better complexity could have been achieved (https://www.geeksforgeeks.org/disjoint-set-data-structures/).