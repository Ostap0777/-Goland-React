import React, { useEffect, useState } from "react";
import { fetchCar, type CreateCarPayload } from "../../../../api/cars";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import Input from "../../../../shared/UI/Input";
import ActionZone from "../../../../shared/ActionZone";
import styles from "./styles.module.scss";

interface CarEditModalProps {
  onClose?: () => void;
  id: number;
}

const initialFormState: Omit<CreateCarPayload, "id"> = {
  make: "",
  model: "",
  year: new Date().getFullYear(),
  price: 0,
  sellerName: "",
  mileage: 0,
  description: "",
  images: [],
};

function CarEditModal({ onClose, id }: CarEditModalProps) {
  const [form, setForm] =
    useState<Omit<CreateCarPayload, "id">>(initialFormState);
  const [files, setFiles] = useState<File[]>([]);
  const queryClient = useQueryClient();

  const { data: car, isLoading } = useQuery({
    queryKey: ["car", id],
    queryFn: () => fetchCar(id),
    enabled: Boolean(id),
  });

  useEffect(() => {
    if (car) {
      setForm({
        make: car.make || "",
        model: car.model || "",
        year: car.year || new Date().getFullYear(),
        price: car.price || 0,
        sellerName: car.sellerName || "",
        mileage: car.mileage || 0,
        description: car.description || "",
        images: car.images || [],
      });
    }
  }, [car]);

  const handleEditCar = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!e.target.value) return;
  };

  return (
    <div className={styles.overlay} onClick={onClose}>
      <div className={styles.modalContent} onClick={(e) => e.stopPropagation()}>
        <h2>Створення / Редагування авто</h2>
        <form onSubmit={handleEditCar}>
          <Input
            name="make"
            value={form.make}
            onChange={handleEditCar}
            placeholder="make"
          />
          <Input
            name="model"
            value={form.model}
            onChange={handleEditCar}
            placeholder="model"
          />
          <Input
            name="year"
            type="number"
            value={form.year}
            onChange={handleEditCar}
            placeholder="year"
          />
          <Input
            name="mileage"
            type="number"
            value={form.mileage}
            onChange={handleEditCar}
            placeholder="mileage"
          />
          <Input
            name="description"
            value={form.description}
            onChange={handleEditCar}
            placeholder="description"
          />
          <Input
            name="price"
            type="number"
            value={form.price}
            onChange={handleEditCar}
            placeholder="price"
          />
          <Input
            name="sellerName"
            value={form.sellerName}
            onChange={handleEditCar}
            placeholder="sellerName"
          />

          <ActionZone
            onFilesChange={(uploadedFiles) => setFiles(uploadedFiles)}
          />

          <button type="submit">Створити</button>
          <button
            type="button"
            onClick={onClose}
            className={styles.closeButton}
          >
            Закрити
          </button>
        </form>
      </div>
    </div>
  );
}

export default CarEditModal;
