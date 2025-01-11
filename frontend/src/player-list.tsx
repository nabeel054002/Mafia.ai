import { Player } from "./types";

// PlayerList Component
interface PlayerListProps {
    players: Player[];
  }
  
  const PlayerList: React.FC<PlayerListProps> = ({ players }) => (
    <div className="PlayerList">
      <h3>Alive Players</h3>
      <ul>
        {players.map((player) => (
          <li key={player.id}>
            {player.id} {player.isAI && '(AI)'}
          </li>
        ))}
      </ul>
    </div>
  );

export default PlayerList;
  