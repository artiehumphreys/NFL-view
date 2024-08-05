import React, { useEffect } from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation, useParams, useNavigate } from "react-router-dom";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import MenuBar from "../components/MenuBar.js";

function InjuryPage() {
  const navigate = useNavigate();
  const { game_id, play_id } = useParams();
  const { push } = useNavigation();
  const location = useLocation();

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);
  return (
    <div className="max-h-screen flex h-screen flex-col">
      <Header></Header>
      <div className="flex flex-1 relative justify-center pt-4 overflow-auto mb-4">
        <div className="w-full max-w-4xl">
          <h1 className="text-center font-medium text-4xl mb-6">
            <span
              onClick={() => navigate(`/games/${game_id}`)}
              className="hover:underline decoration-black cursor-pointer"
            >
              Game {game_id}
            </span>{" "}
            Play {play_id}
          </h1>
        </div>
      </div>
      <MenuBar></MenuBar>
      <Footer></Footer>
    </div>
  );
}

export default InjuryPage;
