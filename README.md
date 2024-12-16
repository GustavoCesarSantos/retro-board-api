usuário N -> usuarios_times <- N times

usuário 1 -> 1 regra de acesso ao time
usuário N <- 1 regra de acesso ao time

users
id
name
email
version
created_at
updated_at

teams
id
name
created_at
updated_at

team_members
id
member_id
team_id
role_id
created_at
updated_at

team_roles
id
name
created_at

boards 1 -> N columns
boards 1 <- 1 columns

columns 1 -> N cards
columns 1 <- 1 cards

cards 1 -> 1 member
cards N <- 1 member

boards
id
team_id
name
active
created_at
updated_at

columns
id
board_id
name
color
position
created_at
updated_at

cards
id
column_id
member_id
text
created_at
updated_at

polls 1 -> N options
polls 1 <- 1 options

options 1 -> N votes
options 1 <- 1 votes

votes 1 -> 1 member
votes 1 <- 1 member

polls
id
team_id
name
created_at
updated_at

options
id
poll_id
text
created_at
updated_at

votes
id
member_id
option_id
created_at
