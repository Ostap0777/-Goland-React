import { createBrowserRouter } from "react-router-dom";
import HomePage from "../pages/HomePage/index";
import Profile from "../pages/Profile";
import MyCars from "../pages/MyCars";

export const router = createBrowserRouter([
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
  {
    path: "*",
    element: <div>404: Сторінку не знайдено</div>,
  },
]);
