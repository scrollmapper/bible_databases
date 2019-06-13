<?php
require_once('BibleReferenceMatch.php');
require_once('BibleReference.php');
require_once('BibleVersion.php');

class BibleSearch {
    const TYPE_TEXT = 'T';
    const TYPE_REFERENCE = 'R';

    /** @var \Database */
    public $db = NULL;

    /**
     * BibleSearch constructor.
     *
     * @param \Database $db
     */
    public function __construct($db) {
        $this->db = $db;
    }

    public function type($searchString) {
        return strstr($searchString, ':') ? self::TYPE_REFERENCE : self::TYPE_TEXT;
    }

    /**
     * @param string $searchString
     * @param string $version
     *
     * @return \BibleReference[]|string
     */
    public function search($searchString, $version) {
        try {
            $type = $this->type($searchString);
            $searchString = trim($searchString);
            if (empty($searchString)) {
                return [];
            }
            elseif ($type === self::TYPE_TEXT) {
                return $this->searchForText($searchString, $version);
            }
            else {
                return $this->searchByReference($searchString, $version);
            }
        }
        catch (Exception $exception) {
            return $exception->getMessage();
        }
    }

    /**
     * @param string $searchString
     * @param string $version
     *
     * @return \BibleReference[]
     */
    private function searchForText($searchString, $version) {
        $references = [];
        $str = "$searchString";
        $query = "SELECT *,
                    MATCH(t) AGAINST(? IN BOOLEAN MODE) AS score,
                    ROUND(((MATCH(t) AGAINST(? IN BOOLEAN MODE) / maxScoreTbl.maxScore) * 100)/20) AS normalized
                    FROM t_kjv,(SELECT MAX(MATCH(t) AGAINST(? IN BOOLEAN MODE)) AS maxScore FROM t_kjv WHERE MATCH(t) AGAINST(? IN BOOLEAN MODE)) AS maxScoreTbl
                    WHERE MATCH(t) AGAINST(? IN BOOLEAN MODE)
                    ORDER BY normalized DESC";
        $rows = $this->db->getRows($query, ['sssss', &$str, &$str, &$str, &$str, &$str]);
        foreach ($rows as $row) {
            $reference = new BibleReference(self::TYPE_TEXT);
            $reference->bookName = $this->lookupBookName($row->b);
            $reference->chapterHuman = $row->c;
            $reference->matches[] = (object)$row;
            $references[] = $reference;
        }
        return $references;
    }

    /**
     * @param integer|string $number
     *
     * @return string|NULL
     */
    private function lookupBookName($number = NULL) {
        return $this->db->getSingleString('SELECT n FROM key_english WHERE b=?', ['s', &$number]);
    }

    /**
     * @param string $searchString
     * @param string $version
     *
     * @return \BibleReference[]
     * @throws \Exception
     */
    private function searchByReference($searchString, $version) {
        $references = $this->parseReferences($searchString);
        foreach ($references as $reference) {
            $query = 'SELECT * FROM ' . $version . ' WHERE ' . $this->verseCondition($reference);
            $reference->matches = $this->db->getRows($query);
        }
        return $references;
    }

    /**
     * @param $searchString
     *
     * @return \BibleReference[]
     * @throws \Exception
     */
    private function parseReferences($searchString) {
        $references = [];
        $searchString = preg_replace('!\s+!', ' ', trim($searchString));

        // Split into reference segments. E.g., Gen 1:1-10,Lev 1:1
        $referenceSearchSegments = explode(',', $searchString);

        if (count($referenceSearchSegments) > 0) {
            foreach ($referenceSearchSegments as $referenceSearchSegment) {
                $reference = new BibleReference(self::TYPE_REFERENCE);

                // Split into book and verse.
                $bookVerseParts = explode(' ', $referenceSearchSegment);
                if (count($bookVerseParts) !== 2) {
                    throw new Exception("Invalid reference: '$referenceSearchSegment''");
                }
                $reference->bookNum = $this->pad($this->lookupBookNumber(@$bookVerseParts[0]), 2);
                $reference->bookName = $this->lookupBookName($reference->bookNum);
                if (empty($reference->bookNum) || empty($reference->bookName)) {
                    throw new Exception("Invalid reference: '$referenceSearchSegment''");
                }

                // Split chapter and verse.
                $verseParts = explode(":", trim($bookVerseParts[1]));
                $reference->chapterHuman = trim(@$verseParts[0]);
                $reference->chapterNum = $this->pad($reference->chapterHuman, 3);
                if (count($verseParts) < 2 || empty($reference->bookNum) || empty($reference->bookName)) {
                    throw new Exception("Invalid reference: '$referenceSearchSegment''");
                }

                // Determine if single or range requested.
                $reference->isRange = (strpos($bookVerseParts[1], '-') !== FALSE || empty($verseParts[1]));
                if ($reference->isRange) {
                    $rangeParts = explode('-', $verseParts[1]);
                    if (count($rangeParts) !== 2) {
                        throw new Exception("Invalid reference range: '$referenceSearchSegment''");
                    }
                    $reference->verseNum = $this->pad(@$rangeParts[0], 3);
                    if (@$rangeParts[1]) {
                        $reference->endVerseNum = $this->pad(@$rangeParts[1], 3);
                    }
                }
                else {
                    $reference->verseNum = $this->pad($verseParts[1], 3);
                }
                $references[] = $reference;
            }
        }
        return $references;
    }

    /**
     * @param string  $input
     * @param integer $maxLen
     *
     * @return string
     */
    public function pad($input, $maxLen) {
        return str_pad($input, $maxLen, '0', STR_PAD_LEFT);
    }

    /**
     * @param string $bookName
     *
     * @return string|NULL
     */
    private function lookupBookNumber($bookName = NULL) {
        return $this->db->getSingleNumber('SELECT b FROM key_abbreviations_english WHERE a=?', ['s', &$bookName]);
    }

    /**
     * @param \BibleReference $reference
     *
     * @return mixed
     */
    private function verseCondition($reference) {
        return $reference->isRange ?
            "id BETWEEN {$reference->bookNum}{$reference->chapterNum}{$reference->verseNum} AND {$reference->bookNum}{$reference->chapterNum}{$reference->endVerseNum} " :
            "id={$reference->bookNum}{$reference->chapterNum}{$reference->verseNum}";
    }

    /**
     * @return \BibleVersion[]
     */
    public function getVersions() {
        return $this->db->getRows('SELECT * FROM bible_version_key ORDER BY version');
    }
}
