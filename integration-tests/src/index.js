// index.js
import "./styles/index.css";
import App from "./components/App";
import React from "react";
import ReactDOM from "react-dom";

export default ReactDOM.render(
    <App />,
    document.getElementById("root") || document.createElement("div")
);
