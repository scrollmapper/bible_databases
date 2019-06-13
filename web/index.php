<?php
require('Database.php');
require('BibleSearch.php');
$db = new Database('localhost', 'bible', 'bible42', 'bible');
$db->connect();

$searchString = trim(@$_POST['searchField']);
$selectedVersion = trim(@$_POST['version']);
if (empty($selectedVersion)) {
    $selectedVersion = 't_kjv';
}
$search = new BibleSearch($db);
$searchType = $search->type($searchString);
$orderCol = $searchType === BibleSearch::TYPE_TEXT? 4 : 1;
$orderDir = $searchType === BibleSearch::TYPE_TEXT? 'desc' : 'asc';
$versions = $search->getVersions();
$references = $search->search($searchString, $selectedVersion);
?>
    <html lang="en">
    <head>
        <title>Bible Search</title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" crossorigin="anonymous">
        <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" crossorigin="anonymous"></script>

        <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.6.3/css/all.css" crossorigin="anonymous">

        <script src="https://code.jquery.com/jquery-3.4.1.min.js" crossorigin="anonymous"></script>
        <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.19/css/jquery.dataTables.css">
        <script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.10.19/js/jquery.dataTables.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/mark.js/8.11.1/jquery.mark.min.js"></script>
        <link rel="stylesheet" href="bible.css">
    </head>
    <body>
    <div class="container-fluid">
        <form action="/" method="POST">
            <div class="row header-bar mb-1 p-1">
                <div class="col-4"><h1>Bible Search</h1></div>
                <div class="col-4">
                    <div class="input-group">
                        <input class="form-control" type="search" value="" placeholder="enter 'book chapter.verse, or text to search for" name="searchField" id="searchField">
                        <span class="input-group-append"><button class="btn btn-outline-secondary" type="button"><i class="fa fa-search"></i></button></span>
                    </div>
                </div>
                <div class="col-4">
                    <select name="version">
                        <?php
                        foreach ($versions as $version) {
                            ?>
                            <option value="<?= $version->table ?>"<?= $selectedVersion == $version->table ? ' selected' : '' ?>><?= $version->version ?></option>
                            <?php
                        }
                        ?>
                    </select>
                </div>
            </div>
        </form>
        <?php
        if (!empty($searchString)) {
            ?>
            <div class="row">
                <div class="col text-center">
                    <h4>Results for: <?= $searchString ?></h4>
                </div>
            </div>
            <?php
        }
        if (is_string($references)) {
            ?>
            <div class="error"><?= $references ?></div>
        <?php
        }
        else {
        ?>
            <table class="table table-striped" id="references" data-order='[[ <?=$orderCol?>, "<?=$orderDir?>" ]]'>
                <thead>
                <tr>
                    <th>Book</th>
                    <th>Chapter</th>
                    <th>Verse</th>
                    <th>Text</th>
                    <?php
                    if ($searchType === BibleSearch::TYPE_TEXT) {
                        ?>
                        <th>Rank</th>
                        <?php
                    }
                    ?>
                </tr>
                </thead>
                <tbody>
                <?php
                /** @var \BibleReference[] $references */
                foreach ($references as $reference) {
                    ?>
                    <?php
                    foreach ($reference->matches as $match) {
                        ?>
                        <tr>
                            <td class="book"><?= $reference->bookName ?></td>
                            <td class="chapter"><?= $reference->chapterHuman ?></td>
                            <td class="verse"><?= $match->v ?></td>
                            <td class="text"><?= $match->t ?></td>
                            <?php
                            if ($searchType === BibleSearch::TYPE_TEXT) {
                                ?>
                                <td class="rank"><?= str_repeat('<i class="fa fa-star green"></i>', $match->normalized) ?></td>
                                <?php
                            }
                            ?>
                        </tr>
                        <?php
                    }
                    ?>
                    <?php
                }
                ?>
                </tbody>
            </table>
            <script>
				$(document).ready(function() {
					window.$('#references').DataTable({paging: true, pageLength: 10, deferRender: true, language: {search: 'filter:'}});
					$('#references').mark('<?= $searchString ?>');
				});
            </script>
            <?php
        }
        ?>
    </div>
    </body>
    </html>
<?php
$db->close();

