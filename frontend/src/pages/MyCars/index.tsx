import React, { useState } from "react";
import Button from "../../shared/UI/Button";
import CarFormModal from "./components/CarFormModal";
import { fetchCars } from "../../api/cars";
import { useQuery } from "@tanstack/react-query";
import CarEditModal from "./components/CarEditModal";

function MyCars() {
  const [openCarModal, setOpenCarModal] = useState<boolean>(false);
  const { data: cars } = useQuery({
    queryKey: ["cars"],
    queryFn: fetchCars,
  });

  const [selectedCarId, setSelectedCarId] = useState<number | null>(null);
  console.log(selectedCarId);
  return (
    <div>
      MyCars
      <Button title={"Add a new cars"} onClick={() => setOpenCarModal(true)} />
      <div>
        {cars?.length == 0 && "Empty list of cars"}
        {cars?.map((car) => (
          <div>
            <p key={car.id}>
              {car.make} {car.model}
            </p>
            <Button
              type="button"
              title="Edit"
              onClick={() => setSelectedCarId(car.id)}
            />
          </div>
        ))}
      </div>
      {openCarModal && <CarFormModal onClose={() => setOpenCarModal(false)} />}
      {selectedCarId !== null && (
        <CarEditModal
          onClose={() => setSelectedCarId(null)}
          id={selectedCarId}
        />
      )}
    </div>
  );
}

export default MyCars;
