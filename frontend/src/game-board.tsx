import PlayerList from "./player-list";

// GameBoard Component
interface GameBoardProps {
    gameState: GameState;
    playerId: string;
    castVote: (voteTarget: string) => void;
  }
  
  const GameBoard: React.FC<GameBoardProps> = ({ gameState, playerId, castVote }) => {
    const alivePlayers = gameState.players.filter((player) => player.isAlive);
  
    return (
      <div className="GameBoard">
        <h2>Game ID: {gameState.id}</h2>
        <h3>Current Phase: {gameState.phase}</h3>
        <PlayerList players={alivePlayers} />
        <VotePanel alivePlayers={alivePlayers} playerId={playerId} castVote={castVote} />
        <GameLog gameLog={gameState.gameLog} />
      </div>
    );
  };

export default GameBoard;
