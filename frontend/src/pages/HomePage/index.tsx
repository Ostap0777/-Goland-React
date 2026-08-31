import { useQuery } from "@tanstack/react-query";
import React from "react";
import { fetchCars } from "../../api/cars";

function HomePage() {
  const {
    data: cars,
    isLoading,
    isError,
    error,
    refetch,
  } = useQuery({
    queryKey: ["makes"],
    queryFn: fetchCars,
  });

  if (isLoading) {
    return <div>Завантаження списку авто...</div>;
  }

  if (isError) {
    return (
      <div>
        <p>
          Помилка завантаження:{" "}
          {error instanceof Error ? error.message : "Невідома помилка"}
        </p>
        <button onClick={() => refetch()}>Спробувати знову</button>
      </div>
    );
  }
  return (
    <div>
      <h2>Оголошення про продаж авто</h2>
      {cars?.length === 0 ? (
        <p>Автомобілів поки немає</p>
      ) : (
        <ul>
          {cars?.map((car) => (
            <li key={car.id}>
              <strong>
                {car.id} {car.model}
              </strong>{" "}
              ({car.name}) — ${car.price}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default HomePage;
