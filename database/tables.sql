create table if not exists Romaji (
    id int(5) primary key,
    value varchar(5),
    difficulty varchar(10),
    date_created timestamp default now(),
    date_updated timestamp default now()
);

create table if not exists HiKa(
    id int(5) primary key,
    value varchar(5),
    family varchar(10),
    romaji_id int(5),
    foreign key (romaji_id) references Romaji(id)
)