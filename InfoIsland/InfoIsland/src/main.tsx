import React, { JSX } from "react";
import ReactDOM from "react-dom/client";
import {App} from "./App";
import "./index.css";
import { ChakraProvider } from "@chakra-ui/react";
import {
  BrowserRouter,
  createBrowserRouter,
  Navigate,
  Route,
  Router,
  RouterProvider,
  Routes,
} from "react-router-dom";
import {Home} from "./home/Home";
import {InsightStreamMain }from "./components/InshitStream/InsightStreamMain";
import {Login }from "./components/Login/Login";
import {RequireAuth, AuthProvider } from "./components/Auth/RequireAuth";

const rootElement = document.getElementById("root");
ReactDOM.createRoot(rootElement!).render(
  <React.StrictMode>
    <AuthProvider>
      <BrowserRouter>
        <ChakraProvider>
          <Routes>
            <Route path="/" element={<App />} />
            <Route path="/login" element={<Login />} />
            <Route
              path="/home"
              element={
                <RequireAuth>
                  <Home />
                </RequireAuth>
              }
            />
            <Route
              path="/insight-stream"
              element={
                <RequireAuth>
                  <InsightStreamMain />
                </RequireAuth>
              }
            />
          </Routes>
        </ChakraProvider>
      </BrowserRouter>
    </AuthProvider>
  </React.StrictMode>
);
