DROP TABLE IF  EXISTS users;

CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


DROP TABLE IF  EXISTS drugs;

CREATE TABLE drugs (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    approved boolean NOT NULL DEFAULT false,
    min_dose integer  NOT NULL DEFAULT 0,
    max_dose integer  NOT NULL DEFAULT 0,
    available_at TIMESTAMP NOT NULL
);

-- Add comments to the columns
COMMENT ON COLUMN drugs.id IS 'Unique identifier for drugs';
COMMENT ON COLUMN drugs.name IS 'Nombre de la droga';
COMMENT ON COLUMN drugs.approved IS 'Indica si está aprobada para uso';
COMMENT ON COLUMN drugs.min_dose IS 'Dosis mínima';
COMMENT ON COLUMN drugs.max_dose IS 'Dosis maxima';
COMMENT ON COLUMN drugs.available_at IS 'Desde esta fecha esta permitida utilizarse';



DROP TABLE IF  EXISTS vaccinations;

CREATE TABLE vaccinations (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    drug_id VARCHAR(255) NOT NULL,
    dose integer  NOT NULL DEFAULT 0,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (drug_id) REFERENCES drugs(id)
);

COMMENT ON COLUMN vaccinations.id IS 'Unique identifier for vaccinations';
COMMENT ON COLUMN vaccinations.name IS 'Nombre de la persona que se vacunara';
COMMENT ON COLUMN vaccinations.drug_id IS 'ID de la droga a vacunar';
COMMENT ON COLUMN vaccinations.dose IS 'Dosis por utilizar en la vacunación';
COMMENT ON COLUMN vaccinations.date IS 'Fecha y hora de la vacunación';


