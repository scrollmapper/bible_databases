<?php
class BibleReference {
    /** @var string */
    public $type = NULL;

    /** @var string */
    public $bookNum = NULL;

    /** @var string */
    public $bookName = NULL;

    /** @var string */
    public $chapterNum = NULL;

    /** @var string */
    public $chapterHuman = NULL;

    /** @var string */
    public $verseNum = 001;

    /** @var string */
    public $endVerseNum = 999;

    /** @var string */
    public $isRange = FALSE;

    /** @var \BibleReferenceMatch[] */
    public $matches = [];

    /**
     * BibleReference constructor.
     *
     * @param string $type
     */
    public function __construct($type) {
        $this->type = $type;
    }
}
