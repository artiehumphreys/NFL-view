import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import StartPage from "./pages/StartPage";
import HomePage from "./pages/HomePage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartPage />} exact />
        <Route path="/home" element={<HomePage />} exact />
      </Routes>
    </Router>
  );
}

export default App;
