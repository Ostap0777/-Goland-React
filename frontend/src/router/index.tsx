import { createBrowserRouter } from "react-router-dom";
import HomePage from "../pages/HomePage/index";
import Profile from "../pages/Profile";
import MyCars from "../pages/MyCars";
import LoginPage from "../pages/Login";
import { ProtectedRoute } from "./ProtectedRoute";
import { PublicRoute } from "./PublicRoute";
import RegisterPage from "../pages/Register";

export const router = createBrowserRouter([
  {
    element: <PublicRoute />,
    children: [
      {
        path: "/login",
        element: <LoginPage />,
      },
      {
        path: "/register",
        element: <RegisterPage />,
      },
    ],
  },

  {
    element: <ProtectedRoute />,
    children: [
      {
        path: "/",
        element: <HomePage />,
      },
      {
        path: "/profile",
        element: <Profile />,
      },
      {
        path: "/mycars",
        element: <MyCars />,
      },
    ],
  },

  {
    path: "*",
    element: <div>404: Сторінку не знайдено</div>,
  },
]);
