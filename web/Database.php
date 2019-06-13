<?php

class Database {
    /** @var string */
    public $server = 'localhost';

    /** @var string */
    public $user = 'bible';

    /** @var string */
    public $password = 'bible';

    /** @var string */
    public $name = 'bible';

    /** @var \mysqli */
    public $connection = NULL;

    /**
     * Database constructor.
     *
     * @param string $server
     * @param string $user
     * @param string $password
     * @param string $name
     */
    public function __construct($server, $user, $password, $name) {
        $this->server = $server;
        $this->user = $user;
        $this->password = $password;
        $this->name = $name;
    }

    /**
     * @param bool $dieOnFailure
     *
     * @return \mysqli
     */
    public function connect($dieOnFailure = TRUE) {
        $this->connection = new mysqli($this->server, $this->user, $this->password, $this->name);
        if ($dieOnFailure) {
            if (!$this->connection) {
                die('Unable to connect to the database.');
            }
            if ($this->connection->connect_error) {
                die("Connect Error ({$this->connection->connect_errno}) {$this->connection->connect_error}).");
            }
        }
        return $this->connection;
    }

    /**
     * @param string         $query
     * @param string|integer $arg
     *
     * @return |null
     */
    public function getSingleNumber($query, $arg) {
        return (int)$this->getSingleString($query, $arg);
    }

    /**
     * @param string   $query
     * @param string[] $args
     *
     * @return |null
     */
    public function getSingleString($query, $args = NULL) {
        $num = NULL;
        $stmt = $this->connection->stmt_init();
        $stmt->prepare($query);
        if (!empty($args)) {
            call_user_func_array([$stmt, 'bind_param'], $args);
        }
        $stmt->execute();
        $result = $stmt->get_result();
        if ($row = $result->fetch_array(MYSQLI_NUM)) {
            $num = $row[0];
        }
        $stmt->close();
        return $num;
    }

    /**
     * @param string   $query
     * @param string[] $args
     *
     * @return array
     */
    public function getRows($query, $args = NULL) {
        $rows = [];
        $stmt = $this->connection->stmt_init();
        $stmt->prepare($query);
        if (!empty($args)) {
            call_user_func_array([$stmt, 'bind_param'], $args);
        }
        $stmt->execute();
        $result = $stmt->get_result();
        if ($result) {
            if ($result->num_rows > 0) {
                while ($row = $result->fetch_assoc()) {
                    $rows[] = (object)$row;
                }
            }
        }
        $stmt->close();
        return $rows;
    }

    public function close() {
        if ($this->connection) {
            $this->connection->close();
        }
    }

    private function buildParams($args) {
        $params = [];
        for ($i = 0; $i < count($args); $i++) {
            $params[] = &$args[$i];
        }
        return $params;
    }
}
