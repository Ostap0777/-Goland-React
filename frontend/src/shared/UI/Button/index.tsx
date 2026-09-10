import React from "react";
import styles from "./styles.module.scss";

interface ButtonProps {
  title: string;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void; // 1. Правильний тип функції (опціональний)
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
}

function Button({
  title,
  onClick,
  type = "button",
  disabled = false,
}: ButtonProps) {
  return (
    <button
      type={type}
      className={styles.buttonBlock}
      onClick={onClick}
      disabled={disabled}
    >
      <span className={styles.buttonContent}>{title}</span>
    </button>
  );
}

export default Button;
