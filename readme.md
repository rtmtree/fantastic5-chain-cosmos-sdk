# Fantastic 5 Bl⚽️ckchain

**Fantastic 5** is a Football Fantasy blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

For those who are new to Football Fantasy, it is a game where you can create your own team with your favorite players and compete with other teams. The performance of the players in the real world will be used to calculate the points of the teams. The team with the highest points will be the winner. More information can be found [here](https://fantasy.premierleague.com/help/rules).

In this blockchain, we downsized the game to 5 players per team. Each goal scored by the player will be 6 points and each assist will be 3 points. Captain of the team will get double points.

## Web Frontend + Demo DApp

Coming soon...

## Get started

To start the blockchain.

```
ignite chain serve
```

To save some alias for the addresses so that we can use them later.

```
export alice=$(fantasfived keys show alice -a)
export bob=$(fantasfived keys show bob -a)
```
In this example, Alice and Bob addresses are `cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv` and `cosmos13ze49rf37h9tq46mydmsllaemhwn0l5jpgsswa` respectively.

#### Verify current Matchweek Id

```
fantasfived query fantasfive show-mw-info
```

This should return something like this:

```
MwInfo:
  nextId: "1"
```

Which means that the current Matchweek Id is 1.

#### Create a Team to participate in the Matchweek

```
fantasfived tx fantasfive create-team Ron Ney Mes Mba Sal --from $alice --gas auto
```

This means that Alice creates a team with the players: Ron, Ney, Mes, Mba, Sal. There will be a prompt to confirm the transaction. Type `y` to confirm.
Let's create some more teams.

```
fantasfived tx fantasfive create-team Ras Ant San Lis Ona --from $alice --gas auto
fantasfived tx fantasfive create-team Son Jam Rom Ric Kul --from $bob --gas auto
fantasfived tx fantasfive create-team Hal Kev Kyl Rub Ede --from $bob --gas auto
```

#### List all teams

```
fantasfived query fantasfive list-stored-team
```

This should return something like this:

```
 pagination:
  next_key: null
  total: "0"
storedTeam:
- captainIndex: "0"
  index: "0"
  mwId: "1"
  owner: cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv
  players: Ron-Ney-Mes-Mba-Sal
  points: "-1"
  rank: "0"
- captainIndex: "0"
  index: "1"
  mwId: "1"
  owner: cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv
  players: Ras-Ant-San-Lis-Ona
  points: "-1"
  rank: "0"
- captainIndex: "0"
  index: "2"
  mwId: "1"
  owner: cosmos13ze49rf37h9tq46mydmsllaemhwn0l5jpgsswa
  players: Son-Jam-Rom-Ric-Kul
  points: "-1"
  rank: "0"
- captainIndex: "0"
  index: "3"
  mwId: "1"
  owner: cosmos13ze49rf37h9tq46mydmsllaemhwn0l5jpgsswa
  players: Hal-Kev-Kyl-Rub-Ede
  points: "-1"
  rank: "0"
```
`-1` Points and `0` Rank mean that the matchweek has not announced the performance of the players yet.

ps. For simplicity, we use the first player as the captain of the team.

#### Announce Players Performance of the Matchweek

```
fantasfived tx fantasfive announce-and-create-next-mw 1 Son-4-9-Kane-1-0-Ric-1-1-Hal-9-3 --from $alice --gas auto
```

This means that Alice (someone trustworthy) announces the performance of the players in the Matchweek 1. The performance is: Son Scored 4 goals, 9 assists, Kane Scored 1 goal, 0 assists, Ric Scored 1 goal, 1 assist, Hal Scored 9 goals, 3 assists. This will also create the next Matchweek.
There will be a prompt to confirm the transaction. Type `y` to confirm.

#### List all teams again

```
fantasfived query fantasfive list-stored-team
```

This should return something like this:

```
pagination:
  next_key: null
  total: "0"
storedTeam:
- captainIndex: "0"
  index: "0"
  mwId: "1"
  owner: cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv
  players: Ron-Ney-Mes-Mba-Sal
  points: "0"
  rank: "4"
- captainIndex: "0"
  index: "1"
  mwId: "1"
  owner: cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv
  players: Ras-Ant-San-Lis-Ona
  points: "0"
  rank: "3"
- captainIndex: "0"
  index: "2"
  mwId: "1"
  owner: cosmos13ze49rf37h9tq46mydmsllaemhwn0l5jpgsswa
  players: Son-Jam-Rom-Ric-Kul
  points: "111"
  rank: "2"
- captainIndex: "0"
  index: "3"
  mwId: "1"
  owner: cosmos13ze49rf37h9tq46mydmsllaemhwn0l5jpgsswa
  players: Hal-Kev-Kyl-Rub-Ede
  points: "126"
  rank: "1"
```

This means that teamId(`index`) 3 is the winner of the Matchweek 1 with 126 points and teamId 2 is the runner-up with 111 points.

#### Verify current Matchweek Id again

```
fantasfived query fantasfive show-mw-info
```

This should return something like this:

```
MwInfo:
  nextId: "2"
```

Which means that the Matchweek Id is currently 2.

The process can be repeated for the next Matchweek and so on.

### Events

For each actions, there will be an event emitted. For example, when Alice announces the performance of the players in the Matchweek 1, there will be an event emitted like this:

```
{
                "type": "match-week-announced",
                "attributes": [
                    {
                        "key": "announcer",
                        "value": "cosmos1anwhl0pavlq0s033gp5n9nq94vytj2ac3ne9fv"
                    },
                    {
                        "key": "mw-id",
                        "value": "1"
                    },
                    {
                        "key": "team-count",
                        "value": "4"
                    },
                    {
                        "key": "winner-team-id",
                        "value": "3"
                    },
                    {
                        "key": "players-performance",
                        "value": "Son-4-9-Kane-1-0-Ric-1-1-Hal-9-3"
                    }
                ]
            }
```


## Known issues

- Players should be represented with integers of PlayerId instead of names in the production level.

## Future Development

- `StoredTeam` should be scaffolded with `list`.
- There should be an `sdk.Msg` for listing teams in a specific Matchweek.
- Reward Distribution should be implemented within `x/fantasfive/keeper/msg_server_announce_and_create_next_mw.go`.

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
