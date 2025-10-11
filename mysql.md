# mysql

## créer une db avec un user du même nom

```sh
NAME="something"
PASS="some password here"

CREATE DATABASE $NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER '$NAME'@'localhost' IDENTIFIED BY $PASS;
GRANT ALL PRIVILEGES ON $NAME.* TO '$NAME'@'localhost';
FLUSH PRIVILEGES;
```
