import React from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation } from "react-router-dom";

function InjuryPage() {
  const { push } = useNavigation();
  const location = useLocation();

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);
  return (
    <div className="max-h-screen flex h-screen flex-col">
      <Header path="/home"></Header>
      <MenuBar></MenuBar>
      <Footer></Footer>
    </div>
  );
}

export default InjuryPage;
