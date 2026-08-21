CREATE TABLE areas (
    area_id INTEGER PRIMARY KEY,

    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(255) NOT NULL UNIQUE,
    is_country BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE competitions (
    competition_id INTEGER PRIMARY KEY,

    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(255) NOT NULL UNIQUE,
    competition_type VARCHAR(255) NOT NULL,

    area_id INTEGER NOT NULL REFERENCES areas(area_id)
);

CREATE TABLE editions (
    competition_id INTEGER NOT NULL REFERENCES competitions(competition_id),
    start_year INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL,

    PRIMARY KEY (competition_id, start_year)
);

CREATE TABLE teams (
    team_id INTEGER PRIMARY KEY,

    full_name VARCHAR(255) NOT NULL UNIQUE,
    short_name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL UNIQUE,
    colors VARCHAR(255),

    stadium VARCHAR(255) NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(area_id)
);

CREATE TABLE edition_teams (
    team_id INTEGER NOT NULL REFERENCES teams(team_id),
    competition_id INTEGER NOT NULL,
    start_year INTEGER NOT NULL,

    PRIMARY KEY (team_id, competition_id, start_year),
    FOREIGN KEY (competition_id, start_year)
        REFERENCES editions(competition_id, start_year)
);

CREATE TABLE matches (
    match_id INTEGER PRIMARY KEY,

    home_team_id INTEGER NOT NULL REFERENCES teams(team_id),
    away_team_id INTEGER NOT NULL REFERENCES teams(team_id),
    competition_id INTEGER NOT NULL,
    start_season_year INTEGER NOT NULL,

    home_goals INTEGER,
    away_goals INTEGER,
    half_time_home_goals INTEGER,
    half_time_away_goals INTEGER,

    status VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    stage VARCHAR(255) NOT NULL,
    matchday INTEGER,

    FOREIGN KEY (competition_id, start_season_year)
        REFERENCES editions(competition_id, start_year),
    CHECK (home_team_id <> away_team_id),
    CHECK (
        (home_goals IS NULL AND away_goals IS NULL)
        OR (home_goals >= 0 AND away_goals >= 0)
    ),
    CHECK (
        (half_time_home_goals IS NULL AND half_time_away_goals IS NULL)
        OR (half_time_home_goals >= 0 AND half_time_away_goals >= 0)
    ),
    CHECK (
        half_time_home_goals IS NULL
        OR half_time_away_goals IS NULL
        OR home_goals IS NULL
        OR away_goals IS NULL
        OR (
            half_time_home_goals <= home_goals
            AND half_time_away_goals <= away_goals
        )
    ),
    CHECK (matchday IS NULL OR matchday > 0)
);

CREATE TABLE goal_scorers (
    goal_scorer_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL REFERENCES teams(team_id),
    competition_id INTEGER NOT NULL,
    start_season_year INTEGER NOT NULL,

    name VARCHAR(255) NOT NULL,
    nationality_area_id INTEGER NOT NULL REFERENCES areas(area_id),

    goals INTEGER NOT NULL,
    assists INTEGER NOT NULL,
    goals_from_penalty INTEGER NOT NULL,

    PRIMARY KEY (goal_scorer_id, competition_id, start_season_year),
    FOREIGN KEY (competition_id, start_season_year)
        REFERENCES editions(competition_id, start_year),
    CHECK (goals >= 0),
    CHECK (assists >= 0),
    CHECK (
        goals_from_penalty >= 0
        AND goals_from_penalty <= goals
    )
);
