import { Player } from "./types";

// VotePanel Component
interface VotePanelProps {
    alivePlayers: Player[];
    playerId: string;
    castVote: (voteTarget: string) => void;
  }
  
  const VotePanel: React.FC<VotePanelProps> = ({ alivePlayers, playerId, castVote }) => (
    <div className="VotePanel">
      <h3>Cast Your Vote</h3>
      <ul>
        {alivePlayers.map((player) => (
          <li key={player.id}>
            <button
              onClick={() => castVote(player.id)}
              disabled={player.id === playerId}
            >
              Vote for {player.id}
            </button>
          </li>
        ))}
      </ul>
    </div>
  );

export default VotePanel;
