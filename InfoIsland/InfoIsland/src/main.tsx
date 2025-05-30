import React, { JSX } from "react";
import ReactDOM from "react-dom/client";
import { App } from "./App";
import "./index.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Home } from "./components/Islands/home/Home";
import { InsightStreamMain } from "./components/Islands/InsightStream/InsightStreamMain";
import { Login } from "./components/Login/Login";
import { RequireAuth, AuthProvider } from "./components/Auth/RequireAuth";
import { MobileHome } from "./components/mobile/MobileHome";
import { AddFeedByMobile } from "./components/Islands/InsightStream/AddFeedByMobile";
import { SummarizeIndexPage } from "./components/mobile/summarize/today";
import { Provider } from "@/components/ui/provider";

const rootElement = document.getElementById("root") as HTMLElement;
ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <AuthProvider>
      <BrowserRouter>
        <Provider>
        <Provider>
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
            <Route path="/mobile" element={<MobileHome />} />
            <Route path="/mobile/register" element={<AddFeedByMobile />} />
            <Route
              path="mobile/summarizeToday"
              element={<SummarizeIndexPage />}
            />
            <Route
              path="mobile/summarizeToday"
              element={<SummarizeIndexPage />}
            />
          </Routes>
        </Provider>
        </Provider>
      </BrowserRouter>
    </AuthProvider>
  </React.StrictMode>,
  </React.StrictMode>,
);
