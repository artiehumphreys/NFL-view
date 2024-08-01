import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import { useParams } from "react-router-dom";
import React, { useEffect, useState } from "react";

function GamePage() {
  const { game_id } = useParams();
  const [injuries, setInjuries] = useState([]);

  useEffect(() => {
    console.log(game_id);
    fetch(`http://localhost:8080/game?game=${game_id}`)
      .then((response) => response.json())
      .then((data) => setInjuries(data));
  }, [game_id]);

  return (
    <div className="max-h-screen flex h-screen flex-col">
      <Header path="/home"></Header>
      <div className="flex flex-1 relative justify-center pt-4 overflow-auto mb-4">
        <div className="w-full max-w-4xl">
          <h1 className="text-center font-medium text-4xl mb-6">
            Game {game_id}
          </h1>
          {injuries.map((injury) => (
            <div
              key={injury.PlayID}
              className="p-4 border rounded mb-4 shadow-lg"
            >
              <h2 className="text-xl font-semibold">
                {injury.FirstName} {injury.LastName}
              </h2>
              <ul className="list-disc pl-5 mt-2">
                <li>
                  <strong>Game:</strong> {injury.Game}
                </li>
                <li>
                  <strong>Type:</strong> {injury.Type}
                </li>
                <li>
                  <strong>Position:</strong> {injury.GamePosition}
                </li>
                <li>
                  <strong>Team:</strong> {injury.Team}
                </li>
                <li>
                  <strong>Jersey Number:</strong> {injury.JerseyNumber}
                </li>
              </ul>
            </div>
          ))}
        </div>
      </div>
      <Footer></Footer>
    </div>
  );
}

export default GamePage;
