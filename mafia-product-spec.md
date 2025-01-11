# AI Mafia - Product Specification

## Product Overview
AI Mafia is a text-based social deduction game where players face off against intelligent AI agents. Unlike traditional Mafia games where players pretend to be mafia, here the AI agents are the actual opponents, trying to blend in while eliminating human players.

## Core Game Experience

### Game Setup
- 6-12 players total
- 2-3 AI agents (number hidden from players)
- Game host announces when room is full
- All players receive role confirmation (Human or AI Agent)
- AI agents privately know each other's identities

### Game Flow

#### 1. Night Phase (60 seconds)
- All players receive "Night has fallen" message
- Humans can only observe (no actions available)
- AI agents privately coordinate their target selection
- System announces elimination result: "X was eliminated in the night"

#### 2. Day Phase (5 minutes)
- Opens with "Day begins" announcement
- Shows current survivor count
- All players can:
  - Chat in public channel
  - Share theories
  - Defend themselves
  - Ask questions
  - Call for votes

#### 3. Voting Phase (60 seconds)
- System announces "Time to vote!"
- All players must vote or abstain
- Live vote count displayed
- Final tally revealed
- Eliminated player's true identity revealed (Human/AI)

### Victory Conditions
- Humans win by eliminating all AI agents
- AI wins by surviving until they equal human numbers
- Game ends immediately when either condition is met

## Player Experience

### Human Player Experience
- Must deduce AI agents through:
  - Chat behavior analysis
  - Voting patterns
  - Response timing
  - Discussion consistency
- Can take notes
- Must collaborate with other humans
- Risk of elimination each night

### Commands & Actions
Players can:
- /vote [player_name] - Cast vote
- /unvote - Remove vote
- /list - See alive players
- /votes - Check current vote count
- /timer - Check phase time remaining

### Game Information Display
Players always see:
- Current phase
- Time remaining
- Alive player count
- Vote counts during voting
- Game history (eliminations)

## Unique Selling Points
1. Real AI Opponents
   - Not scripted NPCs
   - Dynamic conversations
   - Strategic decision making
   - Ability to deceive

2. Pure Social Deduction
   - No special roles/powers
   - Focus on observation
   - Emphasis on discussion
   - Strategy over mechanics

3. Quick Games
   - 15-20 minute sessions
   - Clear progression
   - Immediate feedback
   - Replayability through AI variation

## Future Considerations
- Skill-based matchmaking
- Different AI personality types
- Custom game settings
- Tournament mode
- Replay system
