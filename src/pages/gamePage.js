import Header from "../components/Header.js";
import Footer from "../components/Footer.js";

function GamePage() {
  return (
    <div className="min-h-screen flex flex-col">
      <Header path="/home"></Header>
      <Footer></Footer>
    </div>
  );
}
export default GamePage;
