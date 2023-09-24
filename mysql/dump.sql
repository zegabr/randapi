CREATE TABLE languages (
  language VARCHAR(10) NOT NULL PRIMARY KEY,
  phrase VARCHAR(300) NOT NULL
);

INSERT INTO languages (language, phrase)
VALUES ('go', 'hello from go');
INSERT INTO languages (language, phrase)
VALUES ('python', 'hello from python');
