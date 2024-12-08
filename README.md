usuário N -> usuarios_times <- N times

members
id
name
email
version
createdAt
updatedAt

teams
id
name
createdAt

teamMembers
id
userId
teamId
role
createdAt
updatedAt


boards 1 -> N columns
boards 1 <- 1 columns

columns 1 -> N cards
columns 1 <- 1 cards

cards 1 -> 1 member
cards N <- 1 member

boards
id
teamId
name
active
createdAt
updatedAt

columns
id
boardId
name
color
position
createdAt
updatedAt

cards
id
columnId
memberId
text
createdAt
updatedAt


polls 1 -> N options
polls 1 <- 1 options

options 1 -> N votes
options 1 <- 1 votes

votes 1 -> 1 member
votes 1 <- 1 member

polls
id
teamId
name
createdAt
updatedAt

options
id
pollId
text
createdAt
updatedAt

votes
id
memberId
optionId
createdAt
