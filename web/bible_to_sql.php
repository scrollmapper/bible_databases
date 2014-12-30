<?php

//JOSHUA 1:8-10 to 0601008-0601010

//return book number
class bible_to_sql {
	
	protected $book = null;
	protected $chapter = null;
	protected $verse = null;
	protected $endverse = 999;
	protected $range = FALSE;
	
	public function __construct( /*String*/ $string = null, $range = FALSE) {
				
		//split
		$separatedArray = explode(" ",$string);
		$this->book = $this->bookToNumber($separatedArray[0]);
		
		//split chapter and verse
		$separatedVerse = explode(":",$separatedArray[1]);
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
    
    public function bookToNumber($book) {
		
		switch($book) {
			case "Genesis":
				return "01";
				break;
			case "Exodus":
				return "02";
				break;
			case "Leviticus":
				return "03";
				break;
			case "Numbers":
				return "04";
				break;
			case "Deuteronomy":
				return "05";
				break;
			case "Joshua":
				return "06";
				break;
			case "Judges":
				return "07";
				break;
			case "Ruth":
				return "08";
				break;
			case "1 Samuel":
				return "09";
				break;
			case "2 Samuel":
				return "10";
				break;
			case "1 Kings":
				return "11";
				break;
			case "2 Kings":
				return "12";
				break;
			case "1 Chronicles":
				return "13";
				break;
			case "2 Chronicles":
				return "14";
				break;
			
			case "Ezra":
				return "15";
				break;
			case "Nehemiah":
				return "16";
				break;
			case "Esther":
				return "17";
				break;
			case "Job":
				return "18";
				break;
			case "Psalms":
				return "19";
				break;
			case "Proverbs":
				return "20";
				break;
			case "Ecclesiastes":
				return "21";
				break;
			case "Song of Solomon":
				return "22";
				break;
			case "Isaiah":
				return "23";
				break;
			case "Jeremiah":
				return "24";
				break;
			case "Lamentations":
				return "25";
				break;
			case "Ezekiel":
				return "26";
				break;
			case "Daniel":
				return "27";
				break;
			case "Hosea":
				return "28";
				break;
			case "Joel":
				return "29";
				break;
			case "Amos":
				return "30";
				break;
			case "Obadiah":
				return "31";
				break;
			case "Jonah":
				return "32";
				break;
			case "Micah":
				return "33";
				break;
			case "Nahum":
				return "34";
				break;
			case "Habakkuk":
				return "35";
				break;
			case "Zephaniah":
				return "36";
				break;
			case "Haggai":
				return "37";
				break;
			case "Zechariah":
				return "38";
				break;
			case "Malachi":
				return "39";
				break;
			case "Matthew":
				return "40";
				break;
			case "Mark":
				return "41";
				break;
			case "Luke":
				return "42";
				break;
			case "John":
				return "43";
				break;
			case "Acts":
				return "44";
				break;
			case "Romans":
				return "45";
				break;
			case "1 Corinthians":
				return "46";
				break;
			case "2 Corinthians":
				return "47";
				break;
			case "Galatians":
				return "48";
				break;
			case "Ephesians":
				return "49";
				break;
			case "Philippians":
				return "50";
				break;
			case "Colosians":
				return "51";
				break;
			case "1 Thessalonians":
				return "52";
				break;
			case "2 Thessalonians":
				return "53";
				break;
			case "1 Timothy":
				return "54";
				break;
			case "2 Timothy":
				return "55";
				break;
			case "Titus":
				return "56";
				break;
			case "Philemon":
				return "57";
				break;
			case "Hebrews":
				return "58";
				break;
			case "James":
				return "59";
				break;
			case "1 Peter":
				return "60";
				break;
			case "2 Peter":
				return "61";
				break;
			case "1 John":
				return "62";
				break;
			case "2 John":
				return "63";
				break;
			case "3 John":
				return "64";
				break;
			case "Jude":
				return "65";
				break;
			case "Revelations":
				return "66";
				break;
			default:
				return "01";
				break;
		}
	}
	
	public function sql() {
		if ($this->range) {
			return $this->book.$this->chapter.$this->verse." and ".$this->book.$this->chapter.$this->endverse;
		} else {
			return $this->book.$this->chapter.$this->verse;
		}
	}
    	
	public function getBook() {
		return $this->book;
	}
	
	public function getChapter() {
		return $this-chapter;
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
