import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import StartPage from "./pages/StartPage";
import HomePage from "./pages/HomePage";
import GamePage from "./pages/GamePage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartPage />} exact />
        <Route path="/home" element={<HomePage />} exact />
        <Route path="/games/:game_id" element={<GamePage />} />
      </Routes>
    </Router>
  );
}

export default App;
