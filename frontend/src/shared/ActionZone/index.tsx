import React, { useState, useRef } from "react";
import styles from "./styles.module.scss";

interface ActionZoreProps {
  onFilesChange?: (files: File[]) => void;
}

function ActionZone({ onFilesChange }: ActionZoreProps) {
  const [selectedFiles, setSelectedFiles] = useState<File[]>([]);
  const [previews, setPreview] = useState<string[]>([]);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files) return;
    console.log(e);

    const filesArray = Array.from(e.target.files);
    console.log(filesArray);
    const newPreview = filesArray.map((el) => URL.createObjectURL(el));

    const updatedFiles = [...selectedFiles, ...filesArray];

    setPreview((prev) => [...prev, ...newPreview]);
    console.log(previews);
    setSelectedFiles((prev) => [...prev, ...filesArray]);
    onFilesChange?.(updatedFiles);
  };

  console.log(previews);

  const handleRemove = (index: number) => {
    URL.revokeObjectURL(previews[index]);

    const updatedFiles = selectedFiles.filter((_, i) => i !== index);

    setPreview((prev) => prev.filter((_, i) => i !== index));
    setSelectedFiles((prev) => prev.filter((_, i) => i !== index));
    onFilesChange?.(updatedFiles);
  };

  return (
    <div className={styles.actionZone}>
      <div
        className={styles.dropzone}
        onClick={() => fileInputRef.current?.click()}
      >
        <input
          type="file"
          ref={fileInputRef}
          multiple
          accept="image/*"
          className={styles.fileInput}
          onChange={handleFileChange}
        />
        <div className={styles.uploadIcon}>📷</div>
        <p className={styles.uploadText}>
          Перетягніть фото сюди або <span>оберіть на пристрої</span>
        </p>
      </div>
      <div className={styles.carouselContainer}>
        <div className={styles.carouselTrack}>
          {previews.map((src, index) => (
            <div key={index} className={styles.imageCard}>
              <img src={src} alt={`Preview ${index + 1}`} />

              {index === 0 && <span className={styles.mainBadge}>Головне</span>}

              <button
                type="button"
                className={styles.removeButton}
                onClick={() => handleRemove(index)}
              >
                ✕
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default ActionZone;
