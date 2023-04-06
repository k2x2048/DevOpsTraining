<?php

function dbconnect(){

    $hostname = "mariadb:3306";
    $usernamedb = "cogip-user";
    $passworddb = "cogip-pw";
    $dbname = "cogip";

    $conn = mysqli_connect($hostname, $usernamedb, $passworddb, $dbname) or die ("unable to connect");
    return $conn;
}
?>
