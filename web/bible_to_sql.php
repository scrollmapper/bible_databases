<?php

function convertToNumber($book = NULL, $database = NULL) {
		
	if (!$database) { die('you forgot to specify the database in your bible_to_sql call.'); }
		
	$query = "SELECT B from bible.key_abbreviations_english WHERE A=?";
	
	$stmt = $database->stmt_init();
	$stmt->prepare($query);
	$stmt->bind_param("s", $book);
	$stmt->execute();
        $result = $stmt->get_result();
        if ($row = $result->fetch_array(MYSQLI_NUM)) {
            return $row[0];
        } 
     $stmt->close();
}

function convertToBook($number = NULL, $database = NULL) {
	if (!$database) { die('you forgot to specify the database in your bible_to_sql call.'); }
	
	$query = "SELECT n from bible.key_english WHERE b=?";
	
	$stmt = $database->stmt_init();
	$stmt->prepare($query);
	$stmt->bind_param("s", $number);
	$stmt->execute();
        $result = $stmt->get_result();
        if ($row = $result->fetch_array(MYSQLI_NUM)) {
            return $row[0];
		}
     $stmt->close();
     
}

//JOSHUA 1:8-10 to 0601008-0601010

//return book number
class bible_to_sql {
	
	protected $book = null;
	protected $bookName = null;
	protected $chapter = null;
	protected $chapterHuman = null;
	protected $verse = 001;
	protected $endverse = 999;
	protected $range = FALSE;
	
	public function __construct($string = NULL, $range = FALSE, $database = NULL) {
		
		//places a . between book name and reference. 
		//i.e. changes "Song of Solomon 9:6" to "Song of Solomon.9:6"
		$string = preg_replace('/\s(\S*)$/', '.$1', trim($string)); //trim end for sanitization.
				
		//split
		$separatedArray = explode(".",$string);
		$this->book = $this->addZeros(convertToNumber($separatedArray[0], $database),2);
		$this->bookName = convertToBook($this->book, $database);
		
		//split chapter and verse
		$separatedVerse = explode(":",$separatedArray[1]);
		$this->chapterHuman = $separatedVerse[0];
		$this->chapter = $this->addZeros($separatedVerse[0],3);
		
		//determine if single or range
		if (strpos($string, '-') !== FALSE) {
				$range = TRUE;
			}
		
		if (!$separatedVerse[1]) {
			$range = TRUE;
		}
		
		//set range
		$this->range = $range;
		
		if ($range) {
			$anotherSplit = explode("-",$separatedVerse[1]);
			$this->verse = $this->addZeros($anotherSplit[0],3);
				if ($anotherSplit[1]) {
					$this->endverse = $this->addZeros($anotherSplit[1],3);
				}
		} else {
			$this->verse = $this->addZeros($separatedVerse[1],3);
		}
		
	}
		
	public function addZeros($input,$max) {
		
		$len = strlen($input);
		
		for ($len; $len < $max; $len++) {
			$input = "0".$input;
		}
		
		return $input;
		
	}
	
	public function sql() {
		if ($this->range) {
			return "id BETWEEN ".$this->book.$this->chapter.$this->verse." and ".$this->book.$this->chapter.$this->endverse." ";
		} else {
			return "id='".$this->book.$this->chapter.$this->verse."'";
		}
	}
    	
	public function getBook() {
		return $this->bookName;
	}
	
	public function getChapter() {
		return $this->chapterHuman;
	}
	
	public function getVerse() {
		return $this->verse;
	}
	
	public function getEndVerse() {
		return $this->endverse;
	}
	
	public function getRange() {
		return $this->range;
	}
	
}
