CREATE TABLE areas (
    area_id INTEGER PRIMARY KEY,

    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(255) NOT NULL UNIQUE,
    is_country BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE stadiums (
    stadium_id INTEGER PRIMARY KEY,

    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE competitions (
    competition_id INTEGER PRIMARY KEY,

    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(255) NOT NULL UNIQUE,
    competition_type VARCHAR(255) NOT NULL,

    area_id INTEGER NOT NULL REFERENCES areas(area_id)
);

CREATE TABLE editions (
    edition_id INTEGER PRIMARY KEY,
    
    competition_id INTEGER NOT NULL REFERENCES competitions(competition_id),
    
    start_year INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL,

    UNIQUE (competition_id, start_year)
);

CREATE TABLE teams (
    team_id INTEGER PRIMARY KEY,

    full_name VARCHAR(255) NOT NULL UNIQUE,
    short_name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL UNIQUE,
    colors VARCHAR(255),

    stadium_id INTEGER REFERENCES stadiums(stadium_id),
    area_id INTEGER NOT NULL REFERENCES areas(area_id)
);

CREATE TABLE matches (
    match_id INTEGER PRIMARY KEY,

    home_team_id INTEGER NOT NULL REFERENCES teams(team_id),
    away_team_id INTEGER NOT NULL REFERENCES teams(team_id),
    stadium_id INTEGER REFERENCES stadiums(stadium_id),
    edition_id INTEGER NOT NULL REFERENCES editions(edition_id),

    home_goals INTEGER,
    away_goals INTEGER,
    half_time_home_goals INTEGER,
    half_time_away_goals INTEGER,

    status VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,

    CHECK (home_team_id <> away_team_id),
    CHECK (home_goals IS NULL AND  away_goals IS NULL OR 
        home_goals >= 0 AND away_goals >= 0),
    CHECK (half_time_home_goals IS NULL AND  half_time_away_goals IS NULL OR 
        half_time_home_goals >= 0 AND half_time_away_goals >= 0),
    CHECK (half_time_home_goals <= home_goals AND half_time_away_goals <= away_goals)

);

CREATE TABLE goal_scorers (
    goal_scorer_id INTEGER PRIMARY KEY,

    team_id INTEGER NOT NULL REFERENCES teams(team_id),
    edition_id INTEGER NOT NULL REFERENCES editions(edition_id),

    name VARCHAR(255)  NOT NULL,

    nationality_area_id INTEGER NOT NULL REFERENCES areas(area_id),

    position VARCHAR(255) NOT NULL,

    goals INTEGER NOT NULL,
    assists INTEGER NOT NULL,
    goals_from_penalty INTEGER NOT NULL,

    CHECK (goals >= 0),
    CHECK (assists >= 0),
    CHECK (
        goals_from_penalty >= 0
        AND goals_from_penalty <= goals
    )
);