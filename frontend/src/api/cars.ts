import axios from "axios";

// Описуємо тип об'єкта авто відповідно до вашої Go-моделі
export interface Car {
  id: number;
  brand: string;
  model: string;
  year: number;
  price: number;
}

// Базовий екземпляр axios (з урахуванням порту вашого Go-бекенду)
const api = axios.create({
  baseURL: "http://localhost:8080/api",
});

// Функція для виконання GET-запиту
export const fetchCars = async (): Promise<Car[]> => {
  const { data } = await api.get<Car[]>("/makes");
  return data;
};
