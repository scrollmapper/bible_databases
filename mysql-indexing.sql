USE bible;

ALTER TABLE t_asv
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_bbe
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_dby
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_kjv
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_wbt
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_web
    MODIFY t VARCHAR(1024) NOT NULL;
ALTER TABLE t_ylt
    MODIFY t VARCHAR(1024) NOT NULL;

CREATE INDEX t_asv_t_index ON t_asv(t);
CREATE INDEX t_bbe_t_index ON t_bbe(t);
CREATE INDEX t_dby_t_index ON t_dby(t);
CREATE INDEX t_kjv_t_index ON t_kjv(t);
CREATE INDEX t_wbt_t_index ON t_wbt(t);
CREATE INDEX t_web_t_index ON t_web(t);
CREATE INDEX t_ylt_t_index ON t_ylt(t);

CREATE FULLTEXT INDEX t_asv_t_fulltext ON t_asv(t);
CREATE FULLTEXT INDEX t_bbe_t_fulltext ON t_bbe(t);
CREATE FULLTEXT INDEX t_dby_t_fulltext ON t_dby(t);
CREATE FULLTEXT INDEX t_kjv_t_fulltext ON t_kjv(t);
CREATE FULLTEXT INDEX t_wbt_t_fulltext ON t_wbt(t);
CREATE FULLTEXT INDEX t_web_t_fulltext ON t_web(t);
CREATE FULLTEXT INDEX t_ylt_t_fulltext ON t_ylt(t);
