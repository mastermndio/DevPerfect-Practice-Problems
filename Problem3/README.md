# Problem 3 - Hypothetically, Conditionally, Interatively, Accidentally In Love

**Description:**

Love is kind. Love is blind. Love is confusing and all over the place. In the words of the chainsmokers, who do you love? And who do you like? and who do you hate?

Every emotion has a point value, as follows:

```
!hate: -2
!dislike: -1
!like: 1
!love: 2
```

Everyone has different opinions, so we've gathered all their thoughts together to see who is truly the most loved!

Sum all the emotions people have about different viewers/streamers to determine who is the most loved.

**Input**:

File `input.txt` containing one command per line. 
Each command will be in the following format:
`!emotion: NAME`

**Expected Output:**

Print the name of the most loved viewer/streamer.

-----

**Example Input:**

`input.txt:`

```
!love melkeydev
!hate roxkstar74
!like zanuss
!dislike baldbeardedbuilder
!like arbaya
!like teej_dv
!dislike theprimeagen
!hate beginbot
!love roxkstar74
!hate teej_dv
!love zanuss
!love arbaya
!love at0ta
!hate blackglasses
!love zanuss
!like pimpdaddyslug
!dislike OctagonalRaven 
```

**Example Output:**

`zanuss`