import { IoMdArrowRoundBack, IoMdHome } from "react-icons/io";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useNavigate } from "react-router-dom";

function Header() {
  const navigate = useNavigate();
  const { goBack } = useNavigation();
  return (
    <header className="bg-blue-500 text-white text-center py-4 items-center flex justify-between">
      <IoMdArrowRoundBack
        size={40}
        className="ml-2 cursor-pointer"
        onClick={goBack}
      ></IoMdArrowRoundBack>
      <IoMdHome
        size={40}
        className="ml-2 cursor-pointer"
        onClick={() => navigate("/home")}
      ></IoMdHome>
      <h1
        className="text-5xl font-bold flex-grow text-center cursor-pointer"
        onClick={() => navigate("/home")}
      >
        NFL View
      </h1>
      <div className="w-24"></div>
    </header>
  );
}

export default Header;
