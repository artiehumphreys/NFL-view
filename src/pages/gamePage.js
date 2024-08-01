import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import { useParams } from "react-router-dom";
import React, { useEffect } from "react";

function GamePage() {
  const { game_id } = useParams();
  useEffect(() => {
    fetch(`http://localhost:8080/?game=${game_id}`)
      .then((response) => response.json())
      .then((data) => console.log(data))
      .catch((error) => console.error("Error fetching search tags:", error));
  });
  return (
    <div className="min-h-screen flex flex-col">
      <Header path="/home"></Header>
      <div className="flex flex-1 relative overflow-auto justify-center pt-4">
        <h1 className="text-center font-medium text-4xl">Game {game_id}</h1>
      </div>
      <Footer></Footer>
    </div>
  );
}
export default GamePage;
