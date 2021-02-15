# Tricks

## 1 How to generate (dump) a table to a .sql script
In Mac Terminal,
```
$ sqlite3 <database_file>
> .output output.sql
> .dump <table_name>
> .quit
```
