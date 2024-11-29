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
