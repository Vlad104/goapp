CREATE TABLE IF NOT EXISTS questions (
  "id"          BIGSERIAL                 PRIMARY KEY,
  "userId"      UUID,
  "text"        VARCHAR                   NOT NULL,
  "createdAt"   TIMESTAMP WITH TIME ZONE  DEFAULT now()
);

ALTER TABLE questions
  ADD CONSTRAINT fk_questions_to_users
      FOREIGN KEY("userId") REFERENCES users("id")
        ON DELETE SET NULL
        ON UPDATE NO ACTION;
