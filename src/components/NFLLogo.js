import React from "react";
import { useSpring, animated } from "@react-spring/web";

function NFLLogo() {
  const props = useSpring({
    from: { transform: "translateY(-100%)" },
    to: { transform: "translateY(50vh)" },
    config: { duration: 500 },
  });
  return (
    <div className="logo-container">
      <animated.img
        src="../../public/nfl_logo.jpg"
        style={props}
        alt="NFL Logo"
        className="logo"
      />
    </div>
  );
}

export default NFLLogo;
