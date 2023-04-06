CREATE DATABASE IF NOT EXISTS `wordpress_db`;
CREATE USER 'wordpress_user'@'%' IDENTIFIED BY 'wordpress_pw';
GRANT ALL ON `wordpress_db`.* TO 'wordpress_user'@'%';
