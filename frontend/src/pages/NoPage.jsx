import React from "react";

const style_404 = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    height: "100vh",
  };

const text_style = {
    color: "#ff00ff",
    fontFamily: "Lucida Console",

}
  

export default function NoPage() {
  return (
    <>
      <div className = "screen-center" style = {style_404}>
        <div className="text-center">
          <div className="text-center">
            <h1>4<span>😁</span>4</h1>
          </div>
          <h2 style = {text_style}>Oops! Page Not Found</h2>
          <p>
            Sorry, but the page you are looking for does not exist, has been
            removed, or is temporarily unavailable.
          </p>
          <a href="/">Back to homepage</a>
        </div>
      </div>
    </>
  );
}
