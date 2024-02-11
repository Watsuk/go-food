import { lazy } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navbar from "@/components/layout/Navbar";

const AuthPage = lazy(() => import("@/pages/Auth"));
const DashboardPage = lazy(() => import("@/pages/Dashboard"));

export default function App() {
  return (
    <Router>
      <div className="h-screen w-screen flex flex-col items-center justify-start ">
        <Navbar />
        <Routes>
          <Route path="/" element={<DashboardPage />} />
          <Route path="/auth/*" element={<AuthPage />} />
          <Route path="/dashboard/*" element={<DashboardPage />} />
        </Routes>
      </div>
    </Router>
  );
}
