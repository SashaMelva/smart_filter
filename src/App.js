import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './styles/settings/app.scss';
import {Route, Routes} from "react-router-dom";
import {paths} from "./paths";

function App() {
  return (
      <Routes>
        <Route path={paths.homePage} element={} />

        <Route path={paths.undefined} element={} />
      </Routes>
  );
}

export default App;