<?php

// TEST: /index.php?b=Exodus%2011,Philippians%201:2,Psalms%2019,1%20John%201:9,Song%20Of%20Solomon%209:6&r=0

$mysqli = new mysqli('localhost', 'bible', 'bible', 'bible');

/*
 * This is the "official" OO way to do it,
 * BUT $connect_error was broken until PHP 5.2.9 and 5.3.0.
 */
if ($mysqli->connect_error) {
    die('Connect Error (' . $mysqli->connect_errno . ') '
            . $mysqli->connect_error);
}

echo "Hello Handsome<br />";
require("bible_to_sql.php");
echo "b: ".$_GET['b']." r: ".$_GET['r']."<br />";


//split at commas
$references = explode(",",$_GET['b']);

foreach ($references as $r) {
			
	$test = new bible_to_sql($r, NULL, $mysqli);
	echo "sql: " . $test->sql() . "<br />";

}


$mysqli->close();
?>
