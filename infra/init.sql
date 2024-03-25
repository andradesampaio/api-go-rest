CREATE TABLE IF NOT EXISTS personalidades (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    profissao VARCHAR(255),
    historia TEXT
);

INSERT INTO personalidades (nome, profissao, historia) VALUES ('Tony Stark', 'Homem de Ferro', 'Gênio, bilionário, playboy, filantropo');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Steve Rogers', 'Capitão América', 'Um soldado que ganhou força e agilidade super-humanas por um soro experimental');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Natasha Romanoff', 'Viúva Negra', 'Ex-espiã da KGB que se tornou uma Vingadora');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Bruce Banner', 'Hulk', 'Cientista que se transforma em um monstro verde quando está com raiva');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Thor', 'Deus do Trovão', 'Príncipe de Asgard e membro dos Vingadores');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Peter Parker', 'Homem-Aranha', 'Estudante e fotógrafo que ganhou habilidades de aranha');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('T Challa', 'Pantera Negra', 'Rei de Wakanda e membro dos Vingadores');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Stephen Strange', 'Doutor Estranho', 'Neurocirurgião que se tornou o Mago Supremo');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Carol Danvers', 'Capitã Marvel', 'Ex-piloto da Força Aérea dos EUA que ganhou superpoderes');
INSERT INTO personalidades (nome, profissao, historia) VALUES ('Scott Lang', 'Homem-Formiga', 'Ex-ladrão que usa um traje que lhe permite encolher e controlar formigas');