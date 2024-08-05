import React from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation } from "react-router-dom";

function InjuryPage() {
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
            Game {game_id}, Play {play_id}
          </h1>
        </div>
      </div>
      <MenuBar></MenuBar>
      <Footer></Footer>
    </div>
  );
}

export default InjuryPage;
