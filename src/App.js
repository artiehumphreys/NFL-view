import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import StartPage from "./pages/StartPage.js";
import HomePage from "./pages/HomePage.js";
import GamePage from "./pages/GamePage.js";
import InjuryPage from "./pages/GamePage.js";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartPage />} exact />
        <Route path="/home" element={<HomePage />} exact />
        <Route path="/games/:game_id" element={<GamePage />} />
        <Route path="/games/:game_id/plays/:play_id" element={<InjuryPage />} />
      </Routes>
    </Router>
  );
}

export default App;
