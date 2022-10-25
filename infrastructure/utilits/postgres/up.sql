CREATE TABLE events
(
    id           TEXT PRIMARY KEY,
    aggregate_id TEXT,
    created_at   TEXT,
    user_id      TEXT,
    revision     INT,
    delta_data   JSON,
    event_type   TEXT
);

CREATE TABLE aggregates(
    id           TEXT,
    meta         INT,
    created_at   TEXT
);

ALTER TABLE aggregates
    ADD CONSTRAINT PK_AggregateID
        PRIMARY KEY (id);

ALTER TABLE events
    ADD CONSTRAINT FK_EventAggregateId
        FOREIGN KEY (aggregate_id) REFERENCES aggregates(id);
