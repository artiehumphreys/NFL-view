import React, { useEffect } from "react";
import NFLLogo from "../components/NFLLogo";
import styles from "../css/StartPage.module.css";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation } from "react-router-dom";

function StartPage() {
  const { push } = useNavigation();
  const location = useLocation();

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);
  return (
    <div className={styles.container}>
      <div className={styles.videoBackground}>
        <video autoPlay loop muted className={styles.video}>
          <source src="/NFL_Kickoff_2022.mp4" type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      </div>
      <NFLLogo></NFLLogo>
    </div>
  );
}

export default StartPage;
