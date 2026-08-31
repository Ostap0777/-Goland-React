import { createBrowserRouter } from "react-router-dom";
import HomePage from "../pages/HomePage/index";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "*",
    element: <div>404: Сторінку не знайдено</div>,
  },
]);
