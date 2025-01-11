# AI Mafia: Anonymous Edition
A social deduction game where players face intelligent AI opponents with anonymous voting mechanics.

## Game Overview
Players engage in a battle of wits against AI agents, using social deduction and anonymous voting to eliminate opponents. All votes are private, with only aggregated results shown to maintain strategic depth.

## Core Game Experience

### Setup Phase
- 6-12 players per game
- 2-3 AI agents (exact number hidden)
- Players receive private role assignments
- Game begins when room fills
- Each player gets a unique identifier

### Game Loop

#### 1. Night Phase (45 seconds)
- AI agents select a human target
- Humans cannot take actions
- Result announcement: "Player X was eliminated"
- No information about AI decision process revealed

#### 2. Discussion Phase (4 minutes)
- Open discussion among all players
- Players can:
  - Share theories
  - Make accusations
  - Defend themselves
  - Strategize openly
- AI agents participate in discussion

#### 3. Anonymous Voting Phase (60 seconds)
- Each player votes privately
- Only confirmation of "Vote Cast" shown
- Players cannot see others' votes
- Live aggregate display:
  ```
  Vote Status:
  - 7/8 players have voted
  - Time remaining: 30s
  ```

#### 4. Results Phase (15 seconds)
- Shows only total votes per player:
  ```
  Final Vote Tally:
  Player A: 3 votes
  Player B: 2 votes
  Player C: 2 votes
  Player D: 1 vote
  ```
- Highest vote-getter is eliminated
- Ties result in random elimination
- Eliminated player's role is revealed

### Victory Conditions
- Humans Win: All AI agents eliminated
- AI Wins: Number of AI agents equals surviving humans

## Player Tools & Commands

### Core Commands
- /vote [player_id] - Cast anonymous vote
- /unvote - Cancel your vote
- /list - See alive players
- /tally - View current vote totals
- /time - Check phase timer

### Information Display
Players always see:
- Current phase
- Time remaining
- Alive player count
- Anonymous vote counts
- Game history log

## Unique Features

### Anonymous Voting System
- Votes are completely private
- Only aggregate results shown
- No voting history tracked
- Unable to prove who voted for whom

### Real-time Vote Display
```
Current Vote Distribution:
Total Votes Cast: 6/8
[##########] Player A (3)
[######    ] Player B (2)
[####      ] Player C (1)
[          ] Player D (0)
```

### Game History Log
```
Day 3 Summary:
- 8 players voted
- Player C eliminated (was Human)
- 5 humans and 2 AI remain
```

## Social Features
- Text chat during discussion phase
- Quick response shortcuts
- Basic player profile stats:
  - Games played
  - Win rate
  - Average survival time

## Future Considerations
- Reputation system
- AI personality variants
- Custom game modes
- Spectator mode
- Tournament support