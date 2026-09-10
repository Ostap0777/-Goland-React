import axios from "axios";

// Описуємо тип об'єкта авто відповідно до вашої Go-моделі
export interface FetchCar {
  id: number;
  make: string;
  model: string;
  year: number;
  price: number;
}

export interface CarImageRequest {
  url: string;
  isMain?: boolean;
  order?: number;
}

export interface CreateCarPayload {
  make: string;
  model: string;
  year: number;
  price: number;
  mileage: number;
  description: string;
  sellerName: string;
  images: CarImageRequest[];
}

// Базовий екземпляр axios (з урахуванням порту вашого Go-бекенду)
const api = axios.create({
  baseURL: "http://localhost:8080/api",
});

// Функція для виконання GET-запиту
export const fetchCars = async (): Promise<FetchCar[]> => {
  const { data } = await api.get<FetchCar[]>("/cars");
  return data;
};
// Передаємо carData у параметр і відправляємо його в axios.post
export const AddCars = async (
  carData: Omit<CreateCarPayload, "id"> | CreateCarPayload,
): Promise<CreateCarPayload> => {
  const { data } = await api.post<CreateCarPayload>("/cars", carData);
  return data;
};

export const fetchCar = async (id: number): Promise<FetchCar> => {
  const { data } = await api.get<FetchCar>(`/cars/${id}`);
  return data;
};
