import React, { useState, useEffect } from 'react';
import './App.css';
import GameBoard from './game-board';
import { GameState } from './types';



// App Component
const App: React.FC = () => {
  const [gameId, setGameId] = useState<string | null>(null);
  const [gameState, setGameState] = useState<GameState | null>(null);
  const [playerId, setPlayerId] = useState<string | null>(null);

  const createGame = async () => {
    const response = await fetch('/create');
    const data = await response.json();
    setGameId(data.gameID);
  };

  const joinGame = async () => {
    if (!gameId) return;
    const response = await fetch(`/join?gameID=${gameId}`);
    const data = await response.json();
    setPlayerId(data.playerID);
  };

  const fetchGameState = async () => {
    if (!gameId) return;
    const response = await fetch(`/status?gameID=${gameId}`);
    const data = await response.json();
    setGameState(data);
  };

  const castVote = async (voteTarget: string) => {
    if (!gameId || !playerId) return;
    await fetch(`/vote?gameID=${gameId}&playerID=${playerId}&voteTarget=${voteTarget}`);
    fetchGameState();
  };

  useEffect(() => {
    if (gameId) fetchGameState();
  }, [gameId]);

  return (
    <div className="App">
      <h1>AI Mafia: Anonymous Edition</h1>
      {!gameId && (
        <button onClick={createGame}>Create Game</button>
      )}
      {gameId && !playerId && (
        <div>
          <button onClick={joinGame}>Join Game</button>
        </div>
      )}
      {playerId && gameState && (
        <GameBoard
          gameState={gameState}
          playerId={playerId}
          castVote={castVote}
        />
      )}
    </div>
  );
};


export default App;
