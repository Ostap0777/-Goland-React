import React, { forwardRef, useId } from "react";
import type { InputHTMLAttributes } from "react";
import styles from "./styles.module.scss";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  helperText?: string;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, helperText, className, id, disabled, ...props }, ref) => {
    const generatedId = useId();
    const inputId = id || generatedId;

    return (
      <div
        className={`${styles.inputContainer} ${disabled ? styles.disabled : ""}`}
      >
        {label && (
          <label htmlFor={inputId} className={styles.label}>
            {label}
          </label>
        )}

        <div className={styles.inputWrapper}>
          <input
            ref={ref}
            id={inputId}
            disabled={disabled}
            className={`${styles.input} ${error ? styles.hasError : ""} ${className || ""}`}
            {...props}
          />
        </div>

        {error ? (
          <span className={styles.errorMessage}>{error}</span>
        ) : helperText ? (
          <span className={styles.helperText}>{helperText}</span>
        ) : null}
      </div>
    );
  },
);

Input.displayName = "Input";
export default Input;
