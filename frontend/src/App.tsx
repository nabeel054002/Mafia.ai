import React, { useState, useEffect } from 'react';
import './App.css';
import GameBoard from './game-board';
import { GameState } from './types';

const App: React.FC = () => {
  const [gameState, setGameState] = useState<GameState | null>(null);
  const [playerId, setPlayerId] = useState<string | null>(null);

  // Connect or start a single game instance
  const connectToGame = async () => {
    const response = await fetch('http://localhost:3000/join');
    const data = await response.json();
    setPlayerId(data.playerID);
    setGameState(data.gameState);
  };

  const fetchGameState = async () => {
    const response = await fetch('http://localhost:3000/status');
    const data = await response.json();
    setGameState(data);
  };

  const castVote = async (voteTarget: string) => {
    if (!playerId) return;
    await fetch(`http://localhost:3000/vote?playerID=${playerId}&voteTarget=${voteTarget}`);
    fetchGameState();
  };

  const healthCheck = async () => {
    const response = await fetch('http://localhost:3000/health');
    console.log("Healthcheck resonse", response);
    return response
  }

  // Establish WebSocket connection
  useEffect(() => {
    if (!playerId) return;
    const socket = new WebSocket('ws://localhost:3000/socket');
    socket.onmessage = (event) => {
      const updatedGameState = JSON.parse(event.data);
      setGameState(updatedGameState);
    };
    return () => socket.close();
  }, [playerId]);

  useEffect(() => {
    fetchGameState();
  }, []);

  useEffect(() => {
    healthCheck()
  }, [])

  return (
    <div className="App">
      <h1>AI Mafia: Anonymous Edition</h1>
      {!playerId && (
        <button onClick={connectToGame}>Start/Join Game</button>
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
