import { useQuery } from "@tanstack/react-query";
import React from "react";
import { fetchCars } from "../../api/cars";
import { Link } from "react-router-dom";

// Вкажіть базовий URL вашого сервера статичних файлів/uploads (якщо є)
const UPLOADS_BASE_URL = "http://localhost:8080/uploads/";

function HomePage() {
  const {
    data: cars,
    isLoading,
    isError,
    error,
    refetch,
  } = useQuery({
    queryKey: ["cars"],
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
      <Link to={"/profile"}>Profile</Link>
      <h2>Оголошення про продаж авто</h2>
      {cars?.length === 0 ? (
        <p>Автомобілів поки немає</p>
      ) : (
        <ul style={{ listStyle: "none", padding: 0 }}>
          {cars?.map((car) => {
            // 1. Пошук головного зображення за isMain (camelCase)
            const mainImage =
              car.images?.find((img) => img.isMain) || car.images?.[0];

            // 2. Формування повного URL зображення
            const imageUrl = mainImage?.url
              ? mainImage.url.startsWith("http")
                ? mainImage.url
                : `${UPLOADS_BASE_URL}${mainImage.url}`
              : null;

            return (
              <li
                key={car.id}
                style={{
                  display: "flex",
                  gap: "16px",
                  marginBottom: "16px",
                  alignItems: "center",
                }}
              >
                {imageUrl ? (
                  <img
                    src={imageUrl}
                    alt={`${car.make} ${car.model}`}
                    style={{
                      width: "120px",
                      height: "80px",
                      objectFit: "cover",
                      borderRadius: "6px",
                    }}
                  />
                ) : (
                  <div
                    style={{
                      width: "120px",
                      height: "80px",
                      backgroundColor: "#e0e0e0",
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      borderRadius: "6px",
                      fontSize: "12px",
                      color: "#666",
                    }}
                  >
                    Немає фото
                  </div>
                )}

                <div>
                  <strong>
                    {car.make} {car.model}
                  </strong>{" "}
                  ({car.year}) — ${car.price}
                </div>
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}

export default HomePage;
