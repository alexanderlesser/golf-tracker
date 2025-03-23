CREATE TABLE IF NOT EXISTS clubs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    loft VARCHAR(255) NOT NULL
);

INSERT INTO clubs (name, loft) VALUES
('Driver', '5.5-12.5° (Typically 8-10.5°)'),
('3-Wood', '14-16°'),
('4-Wood', '15-17°'),
('5-Wood', '17-19°'),
('7-Wood', '20-21°'),
('2-Hybrid', '17-18°'),
('3-Hybrid', '19-20°'),
('4-Hybrid', '21-23°'),
('5-Hybrid', '24-26°'),
('2-Iron', '18-19°'),
('3-Iron', '18-21°'),
('4-Iron', '20-24°'),
('5-Iron', '23-27°'),
('6-Iron', '26-31°'),
('7-Iron', '30-35°'),
('8-Iron', '35-39°'),
('9-Iron', '39-44°'),
('PW', '45-47°'),
('GW', '50-52°'),
('SW', '54-58°'),
('LW', '58-64°'),
('Putter', '3.5-4.5°');