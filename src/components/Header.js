import { IoMdArrowRoundBack } from "react-icons/io";
import { useNavigation } from "../contexts/NavigationContext.js";

function Header() {
  const { goBack } = useNavigation();
  return (
    <header className="bg-blue-500 text-white text-center py-4 items-center flex justify-between">
      <IoMdArrowRoundBack
        size={40}
        className="ml-2 cursor-pointer"
        onClick={goBack}
      ></IoMdArrowRoundBack>
      <h1 className="text-5xl font-bold flex-grow text-center">NFL View</h1>
      <div className="w-10"></div>
    </header>
  );
}

export default Header;
