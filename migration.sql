CREATE SCHEMA metrics;

DROP TABLE IF EXISTS metrics.item_metrics;
CREATE TABLE metrics.item_metrics (
  id  UUID DEFAULT uuid_generate_v4(),
  supermetric1 INT,
  supermetric3 INT,
  supermetric2 INT,
  created TIMESTAMP DEFAULT now()
);

-- Use for testing
INSERT INTO metrics.item_metrics(supermetric1, supermetric2, supermetric3)
  SELECT (random()*100 + 20)::int, (random()*100 + 20)::int, (random()*100 + 20)::int
  FROM generate_series (1,200) AS x(n) ;