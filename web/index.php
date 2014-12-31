<?php
// Edit these variables to meet your environment:
$mysql_server = "localhost";
$mysql_username = "bible";
$mysql_password = "bible";
$mysql_db = "bible"; // this is the default table name

$default_text = "John 3:16";

/*** DO NOT EDIT BELOW THIS LINE (Unless you know what you are doing :) ) ***/

$mysqli = new mysqli($mysql_server, $mysql_username, $mysql_password, $mysql_db);

/*
 * This is the "official" OO way to do it,
 * BUT $connect_error was broken until PHP 5.2.9 and 5.3.0.
 */
if ($mysqli->connect_error) {
    die('Connect Error (' . $mysqli->connect_errno . ') '
            . $mysqli->connect_error);
}

require("bible_to_sql.php");
//echo "b: ".$_GET['b']." r: ".$_GET['r']."<br />";


//split at commas
$references = explode(",",$_GET['b']);


?>
<html>
<head>
<title>Bible Search</title>
</head>
<body>
<header>
<form action="index.php" action="GET">
<!-- TODO: Bible dropdown. Defaults to KJV. -->
<label for="b">Reference(s): </label><input type="text" name="b" value="<?php if ($_GET['b']) { echo $_GET['b']; } else { echo $default_text; } ?>" /><input type="submit" value="Search" /><br />

</form>
</header>
<main>
	<?php 
	//return results
	
	foreach ($references as $r) {
				
		$ret = new bible_to_sql($r, NULL, $mysqli);
		//echo "sql query: " . $ret->sql() . "<br />";
		//SELECT * FROM bible.t_kjv WHERE id BETWEEN 01001001 AND 02001005
		$sqlquery = "SELECT * FROM bible.t_kjv WHERE " . $ret->sql();
		$stmt = $mysqli->stmt_init();
		$stmt->prepare($sqlquery);
		$stmt->execute();
        $result = $stmt->get_result();
        if ($result->num_rows > 0) {
			//$row = $result->fetch_array(MYSQLI_NUM);
			//0: ID 1: Book# 2:Chapter 3:Verse 4:Text
			
			print "<article><header><h1>{$ret->getBook()} {$ret->getChapter()}</h1></header>";
			
            while ($row = $result->fetch_row()) {
			 print "<div class=\"versenum\">${row[3]}</div> <div class=\"versetext\">${row[4]}</div><br />";
			}
			print "</article>";
			
        } else {
			print "Did not understand your input.";
		}
		$stmt->close();
	}



	?>
</main>
<footer>
<form action="index.php" action="GET">
<!-- TODO: Bible dropdown. Defaults to KJV. -->
<label for="b">Reference(s): </label><input type="text" name="b" value="<?php if ($_GET['b']) { echo $_GET['b']; } else { echo "John 3:16"; } ?>" /><input type="submit" value="Search" /><br />

</form>
</footer>
</body>
</html>
<?php $mysqli->close(); ?>
