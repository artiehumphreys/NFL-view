import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import { useParams } from "react-router-dom";

function GamePage() {
  const { game_id } = useParams();
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
