CREATE SCHEMA metrics;

DROP TABLE metrics.item_metrics;
CREATE TABLE metrics.item_metrics (
  id  UUID DEFAULT uuid_generate_v4(),
  supermetric1 INT,
  supermetric2 INT,
  created TIMESTAMP DEFAULT now()
);

-- Use for testing 
INSERT INTO metrics.item_metrics(supermetric1, supermetric2)
        SELECT (random()*100 + 20)::int, (random()*100 + 20)::int
        FROM generate_series (1,300) AS x(n) ;

