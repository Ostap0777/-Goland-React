import React, { useState } from "react";
import styles from "./styles.module.scss";
import Input from "../../../../shared/UI/Input";
import ActionZone from "../../../../shared/ActionZone";
import {
  AddCars,
  type CarImageRequest,
  type CreateCarPayload,
} from "../../../../api/cars";

interface CarFormModalProps {
  onClose?: () => void;
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

function CarFormModal({ onClose }: CarFormModalProps) {
  const [form, setForm] =
    useState<Omit<CreateCarPayload, "id">>(initialFormState);
  const [files, setFiles] = useState<File[]>([]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "number" ? Number(value) : value,
    }));
  };
  const handleCreateCar = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const formattedImages: CarImageRequest[] = files.map((file, index) => ({
        url: file.name,
        isMain: index === 0,
        order: index + 1,
      }));

      const payload: CreateCarPayload = {
        make: form.make,
        model: form.model,
        year: Number(form.year),
        price: Number(form.price),
        mileage: Number(form.mileage),
        description: form.description,
        sellerName: form.sellerName,
        images: formattedImages,
      };

      await AddCars(payload);
      onClose?.();
    } catch (error) {
      console.error("Помилка при створенні авто:", error);
    }
  };
  return (
    <div className={styles.overlay} onClick={onClose}>
      <div className={styles.modalContent} onClick={(e) => e.stopPropagation()}>
        <h2>Створення / Редагування авто</h2>
        <form onSubmit={handleCreateCar}>
          <Input
            name="make"
            value={form.make}
            onChange={handleChange}
            placeholder="make"
          />
          <Input
            name="model"
            value={form.model}
            onChange={handleChange}
            placeholder="model"
          />
          <Input
            name="year"
            type="number"
            value={form.year}
            onChange={handleChange}
            placeholder="year"
          />
          <Input
            name="mileage"
            type="number"
            value={form.mileage}
            onChange={handleChange}
            placeholder="mileage"
          />
          <Input
            name="description"
            value={form.description}
            onChange={handleChange}
            placeholder="description"
          />
          <Input
            name="price"
            type="number"
            value={form.price}
            onChange={handleChange}
            placeholder="price"
          />
          <Input
            name="sellerName"
            value={form.sellerName}
            onChange={handleChange}
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

export default CarFormModal;
