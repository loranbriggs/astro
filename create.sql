drop table moons;
drop table planets_habitable;
drop table planets;
drop table small_solar_system_bodies;
drop table stars;
drop table solar_systems;
drop table blackholes;
drop table galaxies;

create table galaxies(
	id 					int not null primary key,
	name 				varchar(20),
	url					varchar(255),
	type				varchar(255)
);

create table solar_systems(
	id					int not null primary key,
	name				varchar(20),
	url					varchar(255),
	parent			int references galaxies(id)
);

create table stars(
	id 					int not null primary key,
	name 				varchar(20),
	url					varchar(255),
	type				varchar(20),
	age					int,
	mass				int,
	luminosity			int,
	tempurature			int,
	parent		  int references solar_systems(id)
);

create table small_solar_system_bodies(
	id					int not null primary key,
	name				varchar(20),
	url					varchar(255),
	type				varchar(20),
	parent			int references stars(id)
);

create table planets(
	id 					int not null primary key,
	name 				varchar(20),
	url					varchar(255),
	type				varchar(20),
	parent			int references stars(id)
);

create table planets_habitable(
	type				varchar(255) not null primary key,
	habitable			varchar(255)
);

create table moons(
	id 					int not null primary key,
	name 				varchar(20),
	url					varchar(255),
	parent			int references planets(id)
);

create table blackholes(
	id 					int not null primary key,
	name 				varchar(20),
	url					varchar(255),
	type				varchar(50),
	angular_momentum	varchar(20),
	charge				varchar(20),
	mass				varchar(50),
	age					varchar(50),
	star_id				varchar(50)
);

insert into galaxies(id, name, type, url) values(1, 'The Milky Way', 'Spiral', 'http://en.wikipedia.org/wiki/Milky_Way');
insert into galaxies(id, name, type, url) values(2, 'Andromeda', 'Spiral', 'http://en.wikipedia.org/wiki/Andromeda_Galaxy');

insert into solar_systems(id, name, parent, url) values(1, 'The Solar System', 1, 'http://en.wikipedia.org/wiki/Solar_System');
insert into solar_systems(id, name, galaxy_id) values(2, 'Crazy Solar System', 2);

insert into stars(id, name, age, tempurature, mass, luminosity, type, url, parent)
values (1, 'Sol', 100000000, 50000, 500000, 100000, 'medium', 'http://en.wikipedia.org/wiki/Sol', 1);

insert into small_solar_system_bodies(id, name, type, parent, url)
values(1, 'Halleys Comet', 'Comet', 1, 'http://en.wikipedia.org/wiki/Halley%27s_Comet');
insert into small_solar_system_bodies(id, name, type, parent, url)
values(2, 'Vesta', 'Asteroid', 1, 'http://en.wikipedia.org/wiki/4_Vesta');

insert into planets_habitable(type, habitable)
values('Rocky', 'No');
insert into planets_habitable(type, habitable)
values('Goldilocks Planet', 'Yes');
insert into planets_habitable(type, habitable)
values('Gas Giant', 'No');

insert into planets(id, name, type, parent, url)
values(1, 'Mercury', 'Rocky', 1, 'http://en.wikipedia.org/wiki/Mercury_(planet)');
insert into planets(id, name, type, parent, url)
values(2, 'Venus', 'Rocky', 1, 'http://en.wikipedia.org/wiki/Venus');
insert into planets(id, name, type, parent, url)
values(3, 'Earth', 'Goldilocks Planet', 1, 'http://en.wikipedia.org/wiki/Earth');
insert into planets(id, name, type, parent, url)
values(4, 'Mars', 'Rocky', 1, 'http://en.wikipedia.org/wiki/Mars');

insert into blackholes(id, name, url, type, angular_momentum, charge, mass, age, star_id) values(1, 'Sagittarius A*', 'http://en.wikipedia.org/wiki/Sagittarius_A*', 'Supermassive Black Hole', 'Unknown', 'Unknown', '(4.31 ± 0.38) × 10 ^ 6 Solar Masses', '12.5 ± 3 billion years', 'Unknown');

insert into moons(id, name, url, parent) values(1, 'The Moon', 'http://en.wikipedia.org/wiki/Moon', 3);
commit;
