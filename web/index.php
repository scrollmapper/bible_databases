<?php

//$mysqli = new mysqli('localhost', 'bible', 'bible', 'bible');

/*
 * This is the "official" OO way to do it,
 * BUT $connect_error was broken until PHP 5.2.9 and 5.3.0.
 */
//if ($mysqli->connect_error) {
//    die('Connect Error (' . $mysqli->connect_errno . ') '
//            . $mysqli->connect_error);
//}

/*
 * Use this instead of $connect_error if you need to ensure
 * compatibility with PHP versions prior to 5.2.9 and 5.3.0.
 */
//if (mysqli_connect_error()) {
//    die('Connect Error (' . mysqli_connect_errno() . ') '
//            . mysqli_connect_error());
//}

//echo 'Success... ' . $mysqli->host_info . "\n";

//$mysqli->close();
echo "Hello Handsome<br />";
require("bible_to_sql.php");
echo "b: ".$_GET['b']." r: ".$_GET['r']."<br />";


//split at commas
$references = explode(",",$_GET['b']);

foreach ($references as $r) {
			
	$test = new bible_to_sql($r, $range);
	echo "sql: " . $test->sql() . "<br />";

}

//$test = new bible_to_sql($_GET['b'],$_GET['r']);
//echo "sql: " . $test->sql();
?>
