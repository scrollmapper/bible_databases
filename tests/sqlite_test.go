package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bookMap map[int]string

func init() {
	db, err := sql.Open("sqlite3", "../bible-sqlite.db?cache=shared&mode=memory")
	check(err)
	bookMap = mapIDToBook(db)
}

func prepareSqliteDB(t *testing.T) (db *sql.DB) {
	db, err := sql.Open("sqlite3", "../bible-sqlite.db?cache=shared&mode=memory")
	assert.Nil(t, err)
	return db
}

type BibleVersionKey struct {
	Table string
}

func returnArray(t *testing.T, query string) []string {
	var result []string
	db := prepareSqliteDB(t)
	rows, err := rowsQuery(query, db)
	assert.Nil(t, err)
	bibleVersionKey := BibleVersionKey{}
	for rows.Next() {
		err = rows.Scan(&bibleVersionKey.Table)
		assert.Nil(t, err)
		result = append(result, bibleVersionKey.Table)
	}
	db.Close()
	return result
}

func TestVerseCount(t *testing.T) {
	expectedCount := 31102
	query := `select "table" from bible_version_key`
	result := returnArray(t, query)
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(*) from %s;", bibleVersion)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in %s translation", expectedCount, bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestBookCount(t *testing.T) {
	expectedCount := 66
	query := `select "table" from bible_version_key`
	result := returnArray(t, query)
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select max(b) from %s;", bibleVersion)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d books in %s translation", expectedCount, bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

// Everything after this line was auto generated using the main program and appended to this file
// go build
// ./tests >> sqlite_test.go

func TestGenesisChapterCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 67
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 55
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 57
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortySixTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFortyNineTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 49
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGenesisChapterFiftyTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 1
	chapterID := 50
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestExodusChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 2
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 59
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 57
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 55
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLeviticusChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 3
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 54
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 89
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 45
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 65
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 54
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 56
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNumbersChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 4
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 68
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDeuteronomyChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 5
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 63
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 45
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoshuaChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 6
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 57
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudgesChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 7
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRuthChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRuthChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 8
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRuthChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 8
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRuthChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 8
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRuthChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 8
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 58
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstSamuelChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 9
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondSamuelChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 10
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 53
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 66
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstKingsChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 53
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 11
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondKingsChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 12
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 54
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 55
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 81
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstChroniclesChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 13
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondChroniclesChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 14
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 70
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzraChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 15
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 73
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNehemiahChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 16
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEstherChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 17
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJobChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 18
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterCount(t *testing.T) {
	expectedCount := 150
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortySixTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFortyNineTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 49
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 50
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyOneTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 51
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 52
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 53
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyFourTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 54
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 55
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftySixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 56
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftySevenTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 57
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyEightTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 58
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterFiftyNineTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 59
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 60
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 61
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 62
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 63
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 64
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 65
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtySixTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 66
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 67
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 68
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSixtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 69
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 70
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyOneTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 71
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 72
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 73
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 74
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 75
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventySixTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 76
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventySevenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 77
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyEightTotalVerseCount(t *testing.T) {
	expectedCount := 72
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 78
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterSeventyNineTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 79
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 80
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyOneTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 81
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 82
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 83
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyFourTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 84
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 85
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightySixTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 86
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightySevenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 87
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyEightTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 88
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterEightyNineTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 89
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 90
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyOneTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 91
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 92
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 93
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 94
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 95
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetySixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 96
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetySevenTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 97
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyEightTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 98
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterNinetyNineTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 99
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 100
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredoneTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 101
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwoTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 102
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 103
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfourTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 104
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfiveTotalVerseCount(t *testing.T) {
	expectedCount := 45
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 105
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredsixTotalVerseCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 106
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredsevenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 107
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredeightTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 108
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundrednineTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 109
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 110
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredelevenTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 111
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwelveTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 112
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirteenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 113
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfourteenTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 114
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfifteenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 115
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredsixteenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 116
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredseventeenTotalVerseCount(t *testing.T) {
	expectedCount := 2
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 117
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredeighteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 118
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundrednineteenTotalVerseCount(t *testing.T) {
	expectedCount := 176
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 119
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 120
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 121
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 122
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 123
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 124
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 125
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 126
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 127
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 128
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredtwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 129
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 130
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 131
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 132
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 133
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 134
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 135
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 136
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 137
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 138
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredthirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 139
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 140
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 141
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 142
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 143
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 144
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 145
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortySixTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 146
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 147
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 148
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfortyNineTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 149
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPsalmsChapterOnehundredfiftyTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 19
	chapterID := 150
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestProverbsChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 20
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEcclesiastesChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 21
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSongOfSolomonChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 22
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterCount(t *testing.T) {
	expectedCount := 66
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortySixTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFortyNineTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 49
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 50
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyOneTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 51
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 52
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 53
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyFourTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 54
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 55
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftySixTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 56
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftySevenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 57
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 58
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterFiftyNineTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 59
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 60
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 61
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 62
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 63
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 64
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 65
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestIsaiahChapterSixtySixTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 23
	chapterID := 66
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortySixTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFortyNineTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 49
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFiftyTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 50
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFiftyOneTotalVerseCount(t *testing.T) {
	expectedCount := 64
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 51
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJeremiahChapterFiftyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 24
	chapterID := 52
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 66
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLamentationsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 25
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 63
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterTwentyNineTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtySixTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtySevenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyEightTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterThirtyNineTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyOneTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortySixTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortySevenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEzekielChapterFortyEightTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 26
	chapterID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 45
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestDanielChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 27
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHoseaChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 28
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoelChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 29
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoelChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 29
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoelChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 29
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJoelChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 29
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestAmosChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 30
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestObadiahChapterCount(t *testing.T) {
	expectedCount := 1
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 31
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestObadiahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 31
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJonahChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 32
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJonahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 32
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJonahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 32
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJonahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 32
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJonahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 32
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterCount(t *testing.T) {
	expectedCount := 7
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMicahChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 33
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNahumChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 34
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNahumChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 34
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNahumChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 34
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestNahumChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 34
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHabakkukChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 35
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHabakkukChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 35
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHabakkukChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 35
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHabakkukChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 35
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZephaniahChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 36
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZephaniahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 36
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZephaniahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 36
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZephaniahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 36
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHaggaiChapterCount(t *testing.T) {
	expectedCount := 2
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 37
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHaggaiChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 37
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHaggaiChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 37
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 9
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestZechariahChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 38
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMalachiChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 39
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMalachiChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 39
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMalachiChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 39
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMalachiChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 39
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMalachiChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 39
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 58
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 46
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 75
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 66
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMatthewChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 40
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 45
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 56
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 72
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestMarkChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 41
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 80
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 49
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 56
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 62
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 54
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 59
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 71
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 56
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestLukeChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 53
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 42
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 51
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 54
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 71
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 53
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 59
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 57
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 50
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJohnChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 43
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 47
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 37
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 42
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 60
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 43
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 48
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 52
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 41
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 38
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyThreeTotalVerseCount(t *testing.T) {
	expectedCount := 35
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 23
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyFourTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 24
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyFiveTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 25
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentySixTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 26
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentySevenTotalVerseCount(t *testing.T) {
	expectedCount := 44
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 27
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestActsChapterTwentyEightTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 44
	chapterID := 28
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 36
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRomansChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 45
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 34
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 58
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstCorinthiansChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 46
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondCorinthiansChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 47
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 31
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestGalatiansChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 48
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 32
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 33
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestEphesiansChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 49
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilippiansChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 50
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilippiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 50
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilippiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 30
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 50
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilippiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 50
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilippiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 50
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestColossiansChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 51
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestColossiansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 51
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestColossiansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 23
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 51
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestColossiansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 51
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestColossiansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 51
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstThessaloniansChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 52
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondThessaloniansChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 53
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondThessaloniansChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 12
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 53
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondThessaloniansChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 53
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondThessaloniansChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 53
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterCount(t *testing.T) {
	expectedCount := 6
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstTimothyChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 54
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondTimothyChapterCount(t *testing.T) {
	expectedCount := 4
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 55
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondTimothyChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 55
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondTimothyChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 55
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondTimothyChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 55
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondTimothyChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 55
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestTitusChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 56
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestTitusChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 56
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestTitusChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 56
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestTitusChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 56
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilemonChapterCount(t *testing.T) {
	expectedCount := 1
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 57
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestPhilemonChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 57
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 16
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 28
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 39
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 40
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestHebrewsChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 58
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 26
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJamesChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 59
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstPeterChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 60
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondPeterChapterCount(t *testing.T) {
	expectedCount := 3
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 61
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondPeterChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 61
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondPeterChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 61
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondPeterChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 61
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterCount(t *testing.T) {
	expectedCount := 5
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 10
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestFirstJohnChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 62
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondJohnChapterCount(t *testing.T) {
	expectedCount := 1
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 63
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestSecondJohnChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 63
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestThirdJohnChapterCount(t *testing.T) {
	expectedCount := 1
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 64
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestThirdJohnChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 64
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudeChapterCount(t *testing.T) {
	expectedCount := 1
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 65
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestJudeChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 25
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 65
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct c) from %s where b = %d;", bibleVersion, bookID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d chapters in %s book in %s translation", expectedCount, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterOneTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 1
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTwoTotalVerseCount(t *testing.T) {
	expectedCount := 29
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 2
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterThreeTotalVerseCount(t *testing.T) {
	expectedCount := 22
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 3
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterFourTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 4
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterFiveTotalVerseCount(t *testing.T) {
	expectedCount := 14
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 5
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterSixTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 6
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterSevenTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 7
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterEightTotalVerseCount(t *testing.T) {
	expectedCount := 13
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 8
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterNineTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 9
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTenTotalVerseCount(t *testing.T) {
	expectedCount := 11
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 10
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterElevenTotalVerseCount(t *testing.T) {
	expectedCount := 19
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 11
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTwelveTotalVerseCount(t *testing.T) {
	expectedCount := 17
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 12
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterThirteenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 13
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterFourteenTotalVerseCount(t *testing.T) {
	expectedCount := 20
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 14
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterFifteenTotalVerseCount(t *testing.T) {
	expectedCount := 8
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 15
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterSixteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 16
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterSeventeenTotalVerseCount(t *testing.T) {
	expectedCount := 18
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 17
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterEighteenTotalVerseCount(t *testing.T) {
	expectedCount := 24
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 18
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterNineteenTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 19
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTwentyTotalVerseCount(t *testing.T) {
	expectedCount := 15
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 20
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTwentyOneTotalVerseCount(t *testing.T) {
	expectedCount := 27
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 21
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}

func TestRevelationChapterTwentyTwoTotalVerseCount(t *testing.T) {
	expectedCount := 21
	query := "select `table` from bible_version_key;"
	result := returnArray(t, query)
	bookID := 66
	chapterID := 22
	for _, bibleVersion := range result {
		query := fmt.Sprintf("select count(distinct v) from %s where b = %d and c = %d;", bibleVersion, bookID, chapterID)
		actualCount, err := returnSingleInt(query)
		assert.Nil(t, err)
		msg := fmt.Sprintf("Should only be %d verses in chapter %d of %s book in %s translation", expectedCount, chapterID, bookMap[bookID], bibleVersion)
		assert.Equal(t, expectedCount, actualCount, msg)
	}
}
