------------------------------------------------------task table--------------------------------------------------------------
CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    done BOOLEAN NOT NULL,
    date TIMESTAMP NOT NULL
);

----------------------------------------------------template table------------------------------------------------------------
CREATE TABLE template (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

---------------------------------------------------template_task table--------------------------------------------------------
CREATE TABLE template_task (
    id SERIAL PRIMARY KEY,
    template_id INTEGER NOT NULL,
    description TEXT NOT NULL,
    FOREIGN KEY (template_id) REFERENCES template (id)
);