// Utility Types
export type Player = {
    id: string;
    role: string;
    isAI: boolean;
    isAlive: boolean;
  };
  
export type GameState = {
    id: string;
    players: Player[];
    phase: string;
    votes: Record<string, number>;
    alivePlayers: number;
    gameLog: string[];
};
