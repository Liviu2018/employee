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


TODO:
    - finish the algorithm
    - read configs from config file
    - add bigger readme
    - add nice UI page (html)
    - add unit tests
    - add delete employee operation

This application exposes a REST API and an html page. The rest endpoints are:
    - /createEmployee  <- will add a new employee into the database, if its manager
    exists and its ID is not already there, and its fields are valid
    - /listAllEmployees <- will return a formated list of all existing employees


Algorithm to format the data O(NlogN):
We arrange the data in a tree, where the CEO is the root, its immediate manages are its children nodes,
and so on. After building* this tree we iterate on it and compose the final result*.

*To build the tree efficiently, we use an auxiliary data structure, namely the parent array. This array
tells us who is the parent of element at index i. parent[i] = parent of employee at index i in out input slice. For any index i, repeatedly calling parent[i], parent[parent[i]], parent[parent[parent[i]]] will
eventually reach the root node (the CEO). If for each employee we thus know the path from him up until the CEO, it is easy to add that employee in a tree.

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