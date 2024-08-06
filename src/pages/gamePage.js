import { useParams, useLocation, useNavigate } from "react-router-dom";
import React, { useEffect, useState } from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import SideBar from "../components/SideBar.js";
import MenuBar from "../components/MenuBar.js";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";

function GamePage() {
  const navigate = useNavigate();

  const { game_id } = useParams();
  const [injuries, setInjuries] = useState([]);
  const [videos, setVideos] = useState([]);

  const { push } = useNavigation();
  const location = useLocation();

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);

  useEffect(() => {
    fetch(`http://localhost:8080/games/${game_id}`)
      .then((response) => response.json())
      .then((data) => {
        setInjuries(data);
      });

    fetch(`http://localhost:8080/videos/broadcast/${game_id}`)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setVideos(data);
      });
  }, [game_id]);

  return (
    <div className="max-h-screen flex h-screen flex-col">
      <Header></Header>
      <div className="flex flex-1 relative justify-center pt-4 overflow-auto mb-4">
        <div className="w-full max-w-4xl">
          <h1 className="text-center font-medium text-4xl mb-6">
            Game {game_id}
          </h1>
          <SideBar></SideBar>
          {injuries.map((injury, index) => (
            <div
              key={index}
              className="p-4 border rounded mb-4 shadow-lg flex-row flex justify-between cursor-pointer"
              onClick={() =>
                navigate(`/games/${injury.Game}/plays/${injury.PlayID}`)
              }
            >
              <div>
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
              {videos[index] && (
                <video
                  controls
                  className="w-1/2"
                  src={`${process.env.PUBLIC_URL}/alpha/nfl_videos/${videos[index]}`}
                >
                  Your browser does not support the video tag.
                </video>
              )}
              <MenuBar></MenuBar>
            </div>
          ))}
        </div>
      </div>
      <Footer></Footer>
    </div>
  );
}

export default GamePage;
