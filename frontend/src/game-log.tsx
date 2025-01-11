// GameLog Component
interface GameLogProps {
    gameLog: string[];
  }
  
  const GameLog: React.FC<GameLogProps> = ({ gameLog }) => (
    <div className="GameLog">
      <h3>Game Log</h3>
      <ul>
        {gameLog.map((entry, index) => (
          <li key={index}>{entry}</li>
        ))}
      </ul>
    </div>
  );

export default GameLog;
