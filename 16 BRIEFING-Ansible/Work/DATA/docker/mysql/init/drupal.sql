CREATE DATABASE IF NOT EXISTS `drupal_db`;
CREATE USER 'drupal_user'@'%' IDENTIFIED BY 'drupal_pw';
GRANT ALL ON `drupal_db`.* TO 'drupal_user'@'%';
