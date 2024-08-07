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

  useEffect(() => {
    fetch(`http://localhost:8080/games/${game_id}/plays/${play_id}/videos`)
      .then((response) => response.json())
      .then((data) => {
        data = data || [];
        console.log(data);
        setVideos(data);
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
        </div>
        <MenuBar></MenuBar>
      </div>
      <Footer></Footer>
    </div>
  );
}

export default PlayPage;
