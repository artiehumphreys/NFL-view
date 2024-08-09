import React, { useEffect, useState } from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation, useParams, useNavigate } from "react-router-dom";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import MenuBar from "../components/MenuBar.js";
import SideBar from "../components/SideBar.js";

function PlayPage() {
  const navigate = useNavigate();
  const { game_id, play_id } = useParams();
  const { push } = useNavigation();
  const location = useLocation();
  const [videos, setVideos] = useState([]);
  const [injury, setInjury] = useState([]);
  const [isConversionSuccessful, setIsConversionSuccessful] = useState(false);

  useEffect(() => {
    fetch(`http://localhost:8080/games/${game_id}/plays/${play_id}/videos`)
      .then((response) => response.json())
      .then((data) => {
        data = data || [];
        setVideos(data);
      });

    fetch(`http://localhost:8080/games/${game_id}/plays/${play_id}`)
      .then((response) => response.json())
      .then((data) => {
        data = data || [];
        setInjury(data);
      });

    fetch(
      `http://localhost:8080/games/${game_id}/plays/${play_id}/conversions`,
      {
        method: "POST",
      }
    ).then((response) => {
      if (response.status === 204) {
        setIsConversionSuccessful(true);
      } else {
        setIsConversionSuccessful(false);
        console.error("Conversion failed or response is not 204.");
      }
    });
  }, [game_id, play_id]);

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);
  return (
    <div className="max-h-screen flex h-screen flex-col">
      <Header></Header>
      <div className="flex flex-1 relative justify-center pt-4 overflow-x-hidden mb-4">
        <SideBar></SideBar>
        <div className="w-full max-w-4xl">
          <h1 className="text-center font-medium text-4xl mb-6">
            <span
              onClick={() => navigate(`/games/${game_id}`)}
              className="hover:underline decoration-blue cursor-pointer text-blue-500"
            >
              Game {game_id}
            </span>{" "}
            Play {play_id}
          </h1>
          <div></div>
          <div className="flex flex-row m-2">
            <ul className="list-disc pl-5 mt-2">
              <li>
                <strong>Name:</strong> {injury.FirstName} {injury.LastName}
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
          <div className="flex flex-row m-2">
            {(videos[0] && (
              <video
                controls
                className="w-1/2 mr-1"
                src={`${process.env.PUBLIC_URL}/alpha/nfl_videos/${videos[0]}`}
              ></video>
            )) || <p className="flex items-center">Failed to Load Video.</p>}
            {(videos[1] && (
              <video
                controls
                className="w-1/2"
                src={`${process.env.PUBLIC_URL}/alpha/nfl_videos/${videos[1]}`}
              ></video>
            )) || <p className="flex items-center">Failed to Load Video.</p>}
          </div>
          <div className="w-full mt-2">
            {isConversionSuccessful && (
              <canvas
                src={`${process.env.PUBLIC_URL}/alpha/final_files/${game_id}_${play_id}`}
              ></canvas>
            )}
          </div>
        </div>
        <MenuBar></MenuBar>
      </div>
      <Footer></Footer>
    </div>
  );
}

export default PlayPage;
